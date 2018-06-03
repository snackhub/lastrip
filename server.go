package lastrip

import (
	"io"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"

	"github.com/snackhub/lastrip/protocol"
	"github.com/snackhub/lastrip/transport"
)

type ServerConfigure struct {
	Addr string
}

type LaStripd struct {
	conn   *transport.ServerConn
	prot   protocol.Protocol
	config *ServerConfigure
}

func NewServer(cfg *ServerConfigure) *LaStripd {
	return &LaStripd{config: cfg, prot: protocol.BinaryProtocol{}}
}

func (server *LaStripd) Serve() {
	conn, err := transport.Serve(server.config.Addr)
	if err != nil {
		panic(err)
	}
	server.conn = conn
	server.acceptloop()
}

func (server *LaStripd) acceptloop() {
	for {
		conn, err := server.conn.Accept()
		if err != nil {
			logrus.Errorf("[Accept] %v", err)
			continue
		}
		go server.handleConn(conn)
	}
}

func (server *LaStripd) handleConn(conn *transport.TCPConn) {
	for {
		var buf [4096]byte
		n, err := server.prot.ReadMessage(conn, buf[:])
		if err != nil {
			if err == io.EOF {
				conn.Close()
				return
			}
		}
		data := buf[:n]
		request, err := server.parseRequest(data)
		if err != nil {
			panic(err)
		}
		logrus.Infof("request: %v\n", request.String())
	}
}

func (server *LaStripd) parseRequest(data []byte) (*protocol.Request, error) {
	var request protocol.Request
	if err := proto.Unmarshal(data, &request); err != nil {
		return nil, err
	}
	return &request, nil
}
