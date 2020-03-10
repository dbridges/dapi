package dapi

import (
	"github.com/rivo/tview"

	"github.com/dbridges/dapi/dapi/widgets"
	"github.com/dbridges/dapi/models"
)

// Dapi stores application wide state nad UI
type Dapi struct {
	app     *tview.Application
	layout  *widgets.Layout
	store   models.Store
	project *models.Project
}

// New creates and returns a new application instance
func New(store models.Store) (*Dapi, error) {
	dapi := Dapi{store: store}

	dapi.app = tview.NewApplication()
	dapi.layout = widgets.NewLayout()
	dapi.app.SetRoot(dapi.layout.Root(), true)

	project, err := dapi.store.Load()
	if err != nil {
		return nil, err
	}
	dapi.project = project

	dapi.layout.SetTitle(dapi.project.Name)

	return &dapi, nil
}

// Run starts the main application instance
func (dapi *Dapi) Run() error {
	return dapi.app.Run()
}
