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

	b, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	s.documents = append(s.documents, b)
	return nil
}

func (s MemoryStore) List() ([]Document, error) {
	docs := make([]Document, 0)

	for _, b := range s.documents {
		var doc Document

		err := json.Unmarshal(b, &doc)
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}
