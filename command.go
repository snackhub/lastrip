package lastrip

import (
	"strings"

	"github.com/snackhub/lastrip/command"
	"github.com/snackhub/lastrip/protocol"
)

type Command interface {
	Initialize([]string) error
	Request() *protocol.Request
}

// NewCommand new a command
func NewCommand(name string) Command {
	switch strings.ToLower(name) {
	case "add":
		return &command.Add{}
	}
	return nil
}
