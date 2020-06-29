package main

type MemoryStore struct {
	documents []Document
}

func (s *MemoryStore) Save(doc Document) error {
	err := doc.Validate()
	if err != nil {
		return err
	}

	s.documents = append(s.documents, doc)
	return nil
}

func (s MemoryStore) List() []Document {
	return s.documents
}
