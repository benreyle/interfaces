package main

import (
	"encoding/json"
	"fmt"
	"time"
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

func main() {
	rg := createRG()

	store := MemoryStore{}

	err := store.Save(rg)
	if err != nil {
		panic(err)
	}

	docs := store.List()

	bytes, err := json.Marshal(docs)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
