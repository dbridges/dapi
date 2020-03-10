package models

import "github.com/dbridges/dapi/util"

// Project represents a single set of API Requests
type Project struct {
	Name     string
	Requests []Request
	Store    Store `json:"-"`
}

// Creates a new project with the given name
func NewProject(name string) (*Project, error) {
	path, err := util.ProjectPath(name)
	if err != nil {
		return nil, err
	}
	store := NewJSONStore(path)
	p := Project{Name: name, Store: store}
	return &p, nil
}

// Save persists the project
func (p *Project) Save() error {
	return p.Store.Save(p)
}
