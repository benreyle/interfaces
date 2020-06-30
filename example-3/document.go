package main

const ContextBrazil string = "brazil"

type Document interface {
	Context() string
	Type() string
	Vouchers() []Voucher
	Validate() error
}
