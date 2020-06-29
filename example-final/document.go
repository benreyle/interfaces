package main

import "reflect"

const ContextBrazil string = "brazil"

var DocumentTypes = map[string]Document{}

type Document interface {
	Context() string
	Type() string
	Vouchers() []Voucher
	Validate() error
}

func NewDocument(key string) (Document, bool) {
	data, exists := DocumentTypes[key]

	if !exists {
		return nil, false
	}

	new := reflect.New(reflect.TypeOf(data).Elem())

	return new.Interface().(Document), true
}
