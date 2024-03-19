package cmd

import (
	"fmt"

	"git.int.86labs.cloud/harrisonde/adele-framework"
)

var Command = &adele.Command{
	Name: "inspire",
	Help: "displays an inspirational quote",
}

func (c *Commands) Inspire() string {
	return fmt.Sprintf("Go build something awesome using %s!", c.App.AppName)
}
