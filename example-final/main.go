package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

func createRG() RG {
	return RG{
		Numero:             "376817884",
		OrgaoEmissor:       "SSP",
		UF:                 "SP",
		DataExpedicao:      time.Now().AddDate(-1, 0, 0),
		NaturalidadeEstado: "SP",
		NaturalidadeCidade: "São Paulo",
		Filiacao1:          "Margie Simpson",
		Filiacao2:          "Homer Simpson",
		Comprovantes: []Voucher{
			{
				Type: "frente",
				File: "/img/teste.jpg",
			},
			{
				Type: "verso",
				File: "/img/teste.jpg",
			},
		},
	}
}

func createCPF() CPF {
	return CPF{
		Numero: "25036156005",
		Comprovantes: []Voucher{
			{
				Type: "frente",
				File: "/img/teste.jpg",
			},
		},
	}
}

func main() {
	rg := createDocumentSpec(createRG())
	cpf := createDocumentSpec(createCPF())

	store := MemoryStore{}

	err := store.Save(rg)
	if err != nil {
		panic(err)
	}

	err = store.Save(cpf)
	if err != nil {
		panic(err)
	}

	docs, err := store.List()
	if err != nil {
		panic(err)
	}

	spew.Dump(docs)
}
