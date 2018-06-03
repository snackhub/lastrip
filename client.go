package lastrip

import (
	"strconv"

	"github.com/golang/protobuf/proto"
	"github.com/snackhub/lastrip/protocol"
	"github.com/snackhub/lastrip/transport"
)

type Client struct {
	conn *transport.TCPConn
	prot protocol.Protocol
}

func Dial(addr string) (*Client, error) {
	conn, err := transport.Dial(addr)
	if err != nil {
		panic(err)
	}
	client := createConn(conn)
	client.prot = protocol.BinaryProtocol{}
	return client, nil
}

func createConn(conn *transport.TCPConn) *Client {
	return &Client{conn: conn}
}

func (c *Client) CreateTask(topic, content string, trigger int64) {
	var request protocol.Request
	request.Command = []byte(topic)
	request.Parameters = append(request.Parameters, []byte(content), []byte(strconv.FormatInt(trigger, 10)))
	request.Version = 1
	data, err := proto.Marshal(&request)
	if err != nil {
		panic(err)
	}
	if err = c.prot.WriteUint16(c.conn, uint16(len(data))); err != nil {
		panic(err)
	}
	c.prot.WriteMessage(c.conn, data)
}

func (c *Client) DoCommand(cmd Command) {
	req := cmd.Request()
	req.Version = 1
	c.SendRequest(req)
}

func (c *Client) SendRequest(req *protocol.Request) error {
	data, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	if err = c.prot.WriteUint16(c.conn, uint16(len(data))); err != nil {
		panic(err)
	}
	c.prot.WriteMessage(c.conn, data)
	return nil
}
