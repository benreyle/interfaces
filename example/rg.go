package main

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type RG struct {
	Numero             string
	OrgaoEmissor       string
	UF                 string
	DataExpedicao      time.Time
	NaturalidadeEstado string
	NaturalidadeCidade string
	Filiacao1          string
	Filiacao2          string
	Comprovantes       []Voucher
}

func (d RG) Context() string {
	return ContextBrazil
}

func (d RG) Type() string {
	return "rg"
}

func (d RG) Vouchers() []Voucher {
	return d.Comprovantes
}

func (d RG) Validate() error {
	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))

	return validation.ValidateStruct(&d,
		validation.Field(&d.Numero, validation.Length(1, 255)),
		validation.Field(&d.OrgaoEmissor, validation.In(OrgaoEmissor...)),
		validation.Field(&d.UF, validation.In(UF...)),
		validation.Field(&d.DataExpedicao, validation.Max(now)),
		validation.Field(&d.NaturalidadeEstado, validation.In(UF...)),
		validation.Field(&d.Filiacao1, validation.Length(1, 255)),
		validation.Field(&d.Filiacao2, validation.Length(1, 255)),
		validation.Field(
			&d.Comprovantes,
			validation.Length(1, 2),
			validation.By(AllowedVouches{
				"frente",
				"verso",
			}.CheckIfAllowed),
		),
	)
}

var OrgaoEmissor = []interface{}{
	"SSP",
	"SSPDC",
	"SESP",
}

var UF = []interface{}{
	"MG",
	"RJ",
	"SP",
}
