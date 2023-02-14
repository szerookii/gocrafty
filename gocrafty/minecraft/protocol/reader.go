package protocol

import (
	"encoding/binary"
	"fmt"
	"math"
)

// Reader is a wrapper around an io.Reader that provides methods to read, it's big endian
type Reader struct {
	Data   []byte
	offset int
}

func (r *Reader) Read(p []byte) (int, error) {
	n := copy(p, r.Data[r.offset:])
	r.offset += n

	return n, nil
}

func (r *Reader) ReadBytes(size int) []byte {
	if r.offset+size > len(r.Data) {
		panic("Out of bounds")
	}

	Data := r.Data[r.offset : r.offset+size]
	r.offset += size

	return Data
}

func (r *Reader) Bool() bool {
	b := r.Byte()

	if b == 0x01 {
		return true
	} else if b == 0x00 {
		return false
	} else {
		panic(fmt.Sprintf("Invalid boolean value: %d", b))
	}
}

func (r *Reader) Byte() byte {
	return r.ReadBytes(1)[0]
}

func (r *Reader) Short(x *int16) {
	*x = int16(binary.BigEndian.Uint16(r.ReadBytes(2)))
}

func (r *Reader) UShort(x *uint16) {
	*x = binary.BigEndian.Uint16(r.ReadBytes(2))
}

func (r *Reader) Int(x *int32) {
	*x = int32(binary.BigEndian.Uint32(r.ReadBytes(4)))
}

func (r *Reader) UInt(x *uint32) {
	*x = binary.BigEndian.Uint32(r.ReadBytes(4))
}

func (r *Reader) Long(x *int64) {
	*x = int64(binary.BigEndian.Uint64(r.ReadBytes(8)))
}

func (r *Reader) ULong(x *uint64) {
	*x = binary.BigEndian.Uint64(r.ReadBytes(8))
}

func (r *Reader) Float(x *float32) {
	*x = math.Float32frombits(binary.BigEndian.Uint32(r.ReadBytes(4)))
}

func (r *Reader) Double(x *float64) {
	*x = math.Float64frombits(binary.BigEndian.Uint64(r.ReadBytes(8)))
}

func (r *Reader) String(x *string) {
	var l int32
	r.VarInt(&l)

	*x = string(r.ReadBytes(int(l)))
}

func (r *Reader) VarInt(x *int32) {
	numRead := 0
	result := int32(0)
	for {
		read := r.Byte()
		value := read & 0b01111111
		result |= int32(value) << (7 * numRead)

		numRead++
		if numRead > 5 {
			panic("Varint is too big")
		}

		if (read & 0b10000000) == 0 {
			*x = result
			return
		}
	}
}

func (r *Reader) VarLong(x *int64) {
	numRead := 0
	result := int64(0)
	for {
		read := r.Byte()
		value := read & 0b01111111
		result |= int64(value) << (7 * numRead)

		numRead++
		if numRead > 10 {
			panic("Varlong is too big")
		}

		if (read & 0b10000000) == 0 {
			*x = result
			return
		}
	}
}
