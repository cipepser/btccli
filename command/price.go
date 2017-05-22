package command

import (
	"strings"
)

type PriceCommand struct {
	Meta
}

func (c *PriceCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *PriceCommand) Synopsis() string {
	return ""
}

func (c *PriceCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
