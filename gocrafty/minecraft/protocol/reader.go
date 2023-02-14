package protocol

import (
	"io"
	"unsafe"
)

type Reader struct {
	r interface {
		io.Reader
		io.ByteReader
	}
}

// NewReader creates a new Reader using the io.ByteReader passed as underlying source to read bytes from.
func NewReader(r interface {
	io.Reader
	io.ByteReader
}) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Bool(x *bool) {
	u, err := r.r.ReadByte()

	if err != nil {
		panic(err)
	}

	*x = *(*bool)(unsafe.Pointer(&u))
}

func (r *Reader) Byte(x *int8) {
	var b uint8

	r.UByte(&b)
	*x = int8(b)
}

func (r *Reader) UByte(x *uint8) {
	var err error
	*x, err = r.r.ReadByte()

	if err != nil {
		panic(err)
	}
}

func (r *Reader) Short(x *int16) {
	b := make([]byte, 2)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*int16)(unsafe.Pointer(&b[0]))
}

func (r *Reader) UShort(x *uint16) {
	b := make([]byte, 2)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*uint16)(unsafe.Pointer(&b[0]))
}

func (r *Reader) Int(x *int32) {
	b := make([]byte, 4)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*int32)(unsafe.Pointer(&b[0]))
}

func (r *Reader) UInt(x *uint32) {
	b := make([]byte, 4)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*uint32)(unsafe.Pointer(&b[0]))
}

func (r *Reader) Long(x *int64) {
	b := make([]byte, 8)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*int64)(unsafe.Pointer(&b[0]))
}

func (r *Reader) ULong(x *uint64) {
	b := make([]byte, 8)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*uint64)(unsafe.Pointer(&b[0]))
}

func (r *Reader) Float(x *float32) {
	b := make([]byte, 4)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*float32)(unsafe.Pointer(&b[0]))
}

func (r *Reader) Double(x *float64) {
	b := make([]byte, 8)

	if _, err := r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*float64)(unsafe.Pointer(&b[0]))
}

func (r *Reader) String(x *string) {
	var (
		length int32
		b      []byte
		err    error
	)

	r.VarInt(&length)
	b = make([]byte, length)

	if _, err = r.r.Read(b); err != nil {
		panic(err)
	}

	*x = *(*string)(unsafe.Pointer(&b))
}

// TODO: Identifier

func (r *Reader) VarInt(x *int32) {
	var (
		b   uint8
		err error
	)

	*x = 0
	for i := uint(0); i < 5; i++ {
		b, err = r.r.ReadByte()

		if err != nil {
			panic(err)
		}

		*x |= int32(b&0x7f) << (7 * i)

		if b&0x80 == 0 {
			return
		}
	}

	panic("VarInt too big")
}

func (r *Reader) VarLong(x *int64) {
	var (
		b   uint8
		err error
	)

	*x = 0
	for i := uint(0); i < 10; i++ {
		b, err = r.r.ReadByte()

		if err != nil {
			panic(err)
		}

		*x |= int64(b&0x7f) << (7 * i)

		if b&0x80 == 0 {
			return
		}
	}

	panic("VarLong too big")
}

func (r *Reader) Read(b []byte) (n int, err error) {
	return r.r.Read(b)
}
