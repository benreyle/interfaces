package main

import "github.com/google/uuid"

type MemoryStore struct {
	documents []DocumentSpec
}

func (s *MemoryStore) Save(doc Document) error {
	err := doc.Validate()
	if err != nil {
		return err
	}

	spec := DocumentSpec{
		ID: uuid.New().String(),
		Context: doc.Context(),
		Type:    doc.Type(),
		Data:    doc,
	}

	s.documents = append(s.documents, spec)
	return nil
}

func (s MemoryStore) List() []DocumentSpec {
	return s.documents
}
