package cli

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Version is set at build time
var Version string

type Opts struct {
	Command string
	Name    string
}

func Parse(args []string) (Opts, error) {
	parser := kingpin.New("dapi", "A terminal API client written in Go.")
	parser.Version(Version)

	newCmd := parser.Command("new", "Create a new dapi project")
	newName := newCmd.Arg("name", "Name of project").Required().String()

	openCmd := parser.Command("open", "Open the speciefied dapi project")
	openName := openCmd.Arg("name", "Name of project").Required().String()

	// If no command is specified default to "run"
	// if len(args) <= 1 {
	// 	args = append(args, "run")
	// }

	cmd, err := parser.Parse(args[1:])
	if err != nil {
		return Opts{}, err
	}

	switch cmd {
	case newCmd.FullCommand():
		return Opts{Command: "new", Name: *newName}, nil
	case openCmd.FullCommand():
		return Opts{Command: "open", Name: *openName}, nil
	}

	return Opts{}, fmt.Errorf("Unknown command '%s'", cmd)
}
