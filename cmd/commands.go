package cmd

import (
	"sync"

	"github.com/harrisonde/adel"
	"github.com/harrisonde/adel/cmd"
)

type Commands struct {
	App *adel.Adel
}

var wg sync.WaitGroup

func (c *Commands) Execute(command string) string {

	switch command {
	case "inspire":
		return c.Inspire()
	case "route":
		return c.List()
	default:
		return c.Help()
	}

}

func (c *Commands) Help() string {

	return cmd.GetHelp()
}
