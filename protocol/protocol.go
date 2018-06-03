package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Protocol

type Protocol interface {
	ReadUint16(reader io.Reader) (uint16, error)
	ReadMessage(reader io.Reader, buf []byte) (int, error)

	WriteUint16(io.Writer, uint16) error
	WriteMessage(io.Writer, []byte) error
}

// LaStripProtocol
// |0       1       |2       3       |
// |length          |data...|

type BinaryProtocol struct {
	Length uint16
	Data   []byte
}

func (p BinaryProtocol) ReadMessage(reader io.Reader, buf []byte) (int, error) {
	length, err := p.ReadUint16(reader)
	if len(buf) < int(length) {
		return 0, fmt.Errorf("len(buf) = %d, length = %d", len(buf), length)
	}
	if err != nil {
		return 0, err
	}
	return io.ReadFull(reader, buf[:length])
}

func (p BinaryProtocol) ReadUint16(reader io.Reader) (uint16, error) {
	var buf [2]byte
	n, err := io.ReadFull(reader, buf[:])
	if err != nil {
		return 0, err
	}
	if n != 2 {
		return 0, fmt.Errorf("ReadUint16: n = %d", n)
	}
	value := binary.LittleEndian.Uint16(buf[:])
	return value, nil
}

func (p BinaryProtocol) WriteUint16(writer io.Writer, value uint16) error {
	var buf [2]byte
	binary.LittleEndian.PutUint16(buf[:2], value)
	n, err := writer.Write(buf[:2])
	if err != nil {
		return err
	}
	if n != 2 {
		return fmt.Errorf("WriteUint16: n = %d", n)
	}
	return nil
}
func (p BinaryProtocol) WriteMessage(writer io.Writer, data []byte) error {
	n, err := writer.Write(data)
	if n != len(data) {
		return fmt.Errorf("WriteUint16: n = %d, len(data) = %d", n, len(data))
	}
	return err
}
