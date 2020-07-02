package main

import (
	"fmt"
	"reflect"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Voucher struct {
	Type string
	File string
}

type AllowedVouches []string

func (a AllowedVouches) CheckIfAllowed(val interface{}) error {
	switch reflect.TypeOf(val).Kind() {
	case reflect.Struct:
		if reflect.TypeOf(Voucher{}) != reflect.TypeOf(val) {
			panic(fmt.Errorf(`parameter is "%T" but should be "%T""`, reflect.TypeOf(val), reflect.TypeOf(Voucher{})))
		}

		errs := make(validation.Errors, 0)

		vouch := reflect.ValueOf(val).Interface().(Voucher)

		if !a.CheckInArray(vouch.Type) {
			errs["type"] = fmt.Errorf(`document does not require file "%s"`, vouch.Type)
		}

		err := validation.Validate(&vouch)
		if err != nil {
			valError := err.(validation.Errors)
			for k, v := range valError {
				errs[k] = v
			}
		}

		if len(errs) > 0 {
			return errs
		}

		return nil

	case reflect.Slice, reflect.Array:
		if reflect.TypeOf([]Voucher{}) != reflect.TypeOf(val) {
			panic(fmt.Errorf(`parameter is "%T" but should be "%T""`, reflect.TypeOf(val), reflect.TypeOf([]Voucher{})))
		}

		err := make(validation.Errors, 0)

		slice := reflect.ValueOf(val).Interface().([]Voucher)

		for index, file := range slice {
			er := a.CheckIfAllowed(file)
			if er != nil {
				i := strconv.Itoa(index)

				err[i] = er
			}
		}

		if len(err) > 0 {
			return err
		}
	}

	return nil
}

func (a AllowedVouches) CheckInArray(t string) bool {
	for _, allow := range a {
		if t == allow {
			return true
		}
	}

	return false
}
