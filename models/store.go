package models

import (
	"encoding/json"
	"io/ioutil"
)

// Store provides an interface to persist projects.
type Store interface {
	Load() (*Project, error)
	Save(*Project) error
}

// JSONStore persists data to a JSON file on disk.
type JSONStore struct {
	path string
}

// NewJSONStore creates a new store that will persist JSON to the supplied path.
func NewJSONStore(path string) *JSONStore {
	return &JSONStore{path: path}
}

// Load reads a project JSON file.
func (store *JSONStore) Load() (*Project, error) {
	var p Project
	data, err := ioutil.ReadFile(store.path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Save persists a project to a JSON file.
func (store *JSONStore) Save(p *Project) error {
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(store.path, data, 0644); err != nil {
		return err
	}
	return nil
}
