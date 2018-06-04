package lastrip

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/snackhub/lastrip/protocol"
)

type Command interface {
	Initialize([]string) error
	Request() *protocol.Request
}

func NewCommand(name string) Command {
	switch strings.ToLower(name) {
	case "add":
		return &CmdAdd{}
	}
	return nil
}

type CmdAdd struct {
	Topic          []byte
	Content        []byte
	RawTriggerTime []byte
	TriggerTime    int64
}

func (c *CmdAdd) Initialize(data []string) error {
	if len(data) != 3 {
		return errors.New("invalid parameters count")
	}
	c.Topic = []byte(data[0])
	c.Content = []byte(data[1])
	value, err := strconv.ParseInt(string(data[2]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid trigger time: (%v:%v)", data[2], err)
	}
	c.RawTriggerTime = []byte(data[2])
	c.TriggerTime = value
	return nil
}

func (c *CmdAdd) Request() *protocol.Request {
	var request protocol.Request
	request.Command = []byte("add")
	request.Parameters = append(request.Parameters, c.Content, c.RawTriggerTime)
	return &request
}
