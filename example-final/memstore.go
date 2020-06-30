package main

import "encoding/json"

type MemoryStore struct {
	documents [][]byte
}

func (s *MemoryStore) Save(doc Document) error {
	err := doc.Validate()
	if err != nil {
		return err
	}

	spec := DocumentSpec{
		Context: doc.Context(),
		Type:    doc.Type(),
		Data:    doc,
	}

	b, err := json.Marshal(spec)
	if err != nil {
		return err
	}

	s.documents = append(s.documents, b)
	return nil
}

func (s MemoryStore) List() ([]DocumentSpec, error) {
	docs := make([]DocumentSpec, 0)

	for _, b := range s.documents {
		var doc DocumentSpec

		err := json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
