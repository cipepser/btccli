package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestPriceCommand_implement(t *testing.T) {
	var _ cli.Command = &PriceCommand{}
}
