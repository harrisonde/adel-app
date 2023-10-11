package cmd

import (
	"regexp"
	"strings"
	"sync"

	"github.com/harrisonde/adel"
	"github.com/harrisonde/adel/cmd"
)

type Commands struct {
	App *adel.Adel
}

var wg sync.WaitGroup
var cmdOptions []string
var optionPattern = "(^--[\\w\\d]{0,}|^-[\\w\\d]{0,})"

func (c *Commands) Execute(arg1, arg2, arg3 string, options []string) string {

	cmdOptions = []string{}

	for _, v := range options {
		isOpt, _ := regexp.MatchString(optionPattern, v)
		if isOpt {
			cmdOptions = append(cmdOptions, v)
		}
	}

	switch arg1 {
	case "inspire":
		return c.Inspire()
	case "route":
		return c.List()
	case "oauth":
		if arg2 == "client" {
			return c.doCreateOauthClient(arg1, arg2, arg3)
		}
	}
	return c.Help()
}

func (c *Commands) Help() string {

	return cmd.GetHelp()
}

func (c *Commands) GetOption(option string) string {
	var o string
	for _, v := range cmdOptions {
		oc := strings.ReplaceAll(v, "-", "")

		switchAcceptArg := strings.Split(oc, "=")
		if len(switchAcceptArg) == 1 {
			if oc == option {
				o = oc
			}
		} else if len(switchAcceptArg) == 2 {
			o = switchAcceptArg[1]
		} else {
			if oc == option {
				o = oc
			}
		}
	}
	return o
}

func (c *Commands) HasOption(option string) bool {
	has := false
	for _, v := range cmdOptions {
		oc := strings.ReplaceAll(v, "-", "")
		switchAcceptArg := strings.Split(oc, "=")
		if len(switchAcceptArg) == 1 {
			if switchAcceptArg[0] == option {
				has = true
			}
		} else if len(switchAcceptArg) == 2 {
			if switchAcceptArg[0] == option {
				has = true
			}
		}
	}
	return has
}
