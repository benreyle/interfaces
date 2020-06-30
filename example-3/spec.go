package main

import "encoding/json"

type DocumentSpec struct {
	Context string
	Type    string
	Data    Document
}

func (s *DocumentSpec) UnmarshalJSON(b []byte) error {
	tmp := struct {
		Context string
		Type    string
		Data    json.RawMessage
	}{}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	s.Context = tmp.Context
	s.Type = tmp.Type

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
	}

	return nil
}
