package cli

import (
	"fmt"

	"github.com/dbridges/dapi/models"
	"github.com/dbridges/dapi/util"
)

func New(name string) {
	p, err := models.NewProject(name)
	util.Must(err)
	util.Must(p.Save())
	fmt.Printf("Created '%s' project\n", p.Name)
}
