package main

import (
	"os"

	"github.com/dbridges/dapi/cli"
)

func main() {
	opts, err := cli.Parse(os.Args)
	if err != nil {
		panic(err)
	}

	switch opts.Command {
	case "new":
		cli.New(opts.Name)
	case "open":
		cli.Open(opts.Name)
	}
}

func initProject() {

}
