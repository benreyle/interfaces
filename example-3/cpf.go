package main

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CPF struct {
	Numero       string
	Comprovantes []Voucher
}

func (d CPF) Context() string {
	return ContextBrazil
}

func (d CPF) Type() string {
	return "rg"
}

func (d CPF) Vouchers() []Voucher {
	return d.Comprovantes
}

func (d CPF) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Numero, validation.Length(1, 255)),
		validation.Field(
			&d.Comprovantes,
			validation.Length(1, 2),
			validation.By(AllowedVouches{
				"frente",
			}.CheckIfAllowed),
		),
	)
}
