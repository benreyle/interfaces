package main

import (
	"encoding/json"
	"errors"
)

type DocumentSpec struct {
	Id      string
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

	data, exists := NewDocument(s.Type)
	if !exists {
		return errors.New("invalid document type")
	}

	err = json.Unmarshal(tmp.Data, &data)
	if err != nil {
		return err
	}

	s.Data = data

	return nil
}
