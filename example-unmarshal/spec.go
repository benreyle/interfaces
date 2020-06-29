package main

import (
	"encoding/json"
	"errors"
)

type DocumentSpec struct {
	ID      string
	Context string
	Type    string
	Data    Document
}

func (s *DocumentSpec) UnmarshalJSON(b []byte) error {
	type Alias DocumentSpec

	tmp := struct {
		Alias
		Data json.RawMessage
	}{}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = DocumentSpec(tmp.Alias)

	switch s.Type {
	case "rg":
		data := RG{}

		err = json.Unmarshal(tmp.Data, &data)
		if err != nil {
			return err
		}

		s.Data = data
	case "cpf":
		data := CPF{}

		err = json.Unmarshal(tmp.Data, &data)
		if err != nil {
			return err
		}

		s.Data = data
	default:
		return errors.New("invalid document type")
	}

	return nil
}
