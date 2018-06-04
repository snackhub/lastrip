package transport

import "net"

type TCPConn struct {
	net.Conn
}

func Dial(address string) (*TCPConn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPConn{
		Conn: conn,
	}, nil
}

type ServerConn struct {
	ln net.Listener
}

func Serve(address string) (*ServerConn, error) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &ServerConn{
		ln: ln,
	}, nil
}

func (s *ServerConn) Accept() (*TCPConn, error) {
	conn, err := s.ln.Accept()
	if err != nil {
		return nil, err
	}
	return &TCPConn{Conn: conn}, nil
}
