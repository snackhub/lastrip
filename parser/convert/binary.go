package convert

import (
	"encoding/binary"
	"math"
)

type ByteConverter interface {
	binary.ByteOrder
	Float32([]byte) float32
	Float64([]byte) float64
	PutFloat32([]byte, float32)
	PutFloat64([]byte, float64)
}

type littleEndian struct {
	binary.ByteOrder
}

type bigEndian struct {
	binary.ByteOrder
}

// LittleEndian is the little-endian implementation of ByteConverter.
var LittleEndian = littleEndian{
	ByteOrder: binary.LittleEndian,
}

// BigEndian is the big-endian implementation of ByteConverter.
var BigEndian = bigEndian{
	ByteOrder: binary.BigEndian,
}

// Float32 parse bytes data to a float32 value
func (littleEndian) Float32(b []byte) float32 {
	// bytes to uint32
	f := binary.LittleEndian.Uint32(b)
	// int32 to float32
	return math.Float32frombits(f)
}
func (littleEndian) Float64(b []byte) float64 {
	// bytes to uint32
	f := binary.LittleEndian.Uint64(b)
	// int32 to float32
	return math.Float64frombits(f)
}
func (littleEndian) PutFloat32(b []byte, f float32) {
	// float32 to uint32
	bits := math.Float32bits(f)
	// uint32 to bytes
	binary.LittleEndian.PutUint32(b, bits)
}
func (littleEndian) PutFloat64(b []byte, f float64) {
	// float32 to uint32
	bits := math.Float64bits(f)
	// uint32 to bytes
	binary.LittleEndian.PutUint64(b, bits)
}

func (bigEndian) Float32(b []byte) float32 {
	// bytes to uint32
	f := binary.BigEndian.Uint32(b)
	// int32 to float32
	return math.Float32frombits(f)
}
func (bigEndian) Float64(b []byte) float64 {
	// bytes to uint32
	f := binary.BigEndian.Uint64(b)
	// int32 to float32
	return math.Float64frombits(f)
}
func (bigEndian) PutFloat32(b []byte, f float32) {
	// float32 to uint32
	bits := math.Float32bits(f)
	// uint32 to bytes
	binary.BigEndian.PutUint32(b, bits)

}
func (bigEndian) PutFloat64(b []byte, f float64) {
	// float32 to uint32
	bits := math.Float64bits(f)
	// uint32 to bytes
	binary.BigEndian.PutUint64(b, bits)
}
