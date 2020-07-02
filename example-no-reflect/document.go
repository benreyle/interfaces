package main

const ContextBrazil string = "brazil"

var DocumentTypes = map[string]func() Document{}

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
	return data(), true
}
