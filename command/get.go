package command

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/snackhub/lastrip/protocol"
)

// Get command for getting tasks
type Get struct {
	// Topic target topic
	Topic []byte
	// TriggerTime task trigger time
	TriggerTime int64
	// Count tasks count
	Count int32
}

func (c *Get) Initialize(data []string) error {
	if len(data) != 3 {
		return errors.New("invalid parameters count")
	}
	c.Topic = []byte(data[0])

	value, err := strconv.ParseInt(string(data[1]), 10, 64)
	if err != nil {
		return fmt.Errorf("invalid trigger time: (%v:%v)", data[2], err)
	}
	c.TriggerTime = value
	return nil
}

func (c *Get) Request() *protocol.Request {
	var request protocol.Request
	request.Command = []byte("add")
	request.Parameters = append(request.Parameters, c.TriggerTime, c.Count)
	return &request
}
