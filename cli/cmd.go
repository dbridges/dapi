package cli

import (
	"fmt"
	"os"

	"github.com/dbridges/dapi/dapi"
	"github.com/dbridges/dapi/models"
	"github.com/dbridges/dapi/util"
)

func New(name string) {
	p, err := models.NewProject(name)
	util.Must(err)
	util.Must(p.Save())
	fmt.Printf("Created '%s' project\n", p.Name)
}

func Open(name string) {
	os.Setenv("TCELL_TRUECOLOR", "disable")
	path, err := util.ProjectPath(name)
	util.Must(err)
	store := models.NewJSONStore(path)
	dapi, err := dapi.New(store)
	util.Must(err)
	if err := dapi.Run(); err != nil {
		panic(err)
	}
}
