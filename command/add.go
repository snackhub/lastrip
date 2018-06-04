package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/snackhub/lastrip/protocol"
)

type Add struct {
	Topic       []byte
	Content     []byte
	TriggerTime []byte
}

func (c *Add) Initialize(data []string) error {
	if len(data) != 3 {
		return errors.New("invalid parameters count")
	}
	c.Topic = []byte(data[0])
	c.Content = []byte(data[1])
	value, err := strconv.ParseInt(string(data[2]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid trigger time: (%v:%v)", data[2], err)
	}
	c.TriggerTime = []byte(data[2])
	return nil
}

func (c *Add) Request() *protocol.Request {
	var request protocol.Request
	request.Command = []byte("add")
	request.Parameters = append(request.Parameters, c.Content, c.TriggerTime)
	return &request
}
