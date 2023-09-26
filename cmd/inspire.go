package cmd

import (
	"fmt"

	"github.com/harrisonde/adel"
)

var Command = &adel.Command{
	Name: "inspire",
	Help: "displays an inspirational quote",
}

func (c *Commands) Inspire() string {
	return fmt.Sprintf("Go build something awesome using %s!", c.App.AppName)
}
