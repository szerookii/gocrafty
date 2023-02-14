package protocol

import (
	"io"
	"unsafe"
)

type Writer struct {
	w interface {
		io.Writer
		io.ByteWriter
	}
}

func NewWriter(w interface {
	io.Writer
	io.ByteWriter
}) *Writer {
	return &Writer{w: w}
}

func (w *Writer) Bool(x bool) {
	_ = w.w.WriteByte(*(*byte)(unsafe.Pointer(&x)))
}

func (w *Writer) Byte(x int8) {
	w.w.WriteByte(byte(x) & 0xff)
}

func (w *Writer) UByte(x uint8) {
	w.w.WriteByte(x)
}

func (w *Writer) Short(x int16) {
	data := *(*[2]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) UShort(x uint16) {
	data := *(*[2]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) Int(x int32) {
	data := *(*[4]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) UInt(x uint32) {
	data := *(*[4]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) Long(x int64) {
	data := *(*[8]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) ULong(x uint64) {
	data := *(*[8]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) Float(x float32) {
	data := *(*[4]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) Double(x float64) {
	data := *(*[8]byte)(unsafe.Pointer(&x))
	_, _ = w.w.Write(data[:])
}

func (w *Writer) String(x string) {
	w.VarInt(int32(len(x)))
	_, _ = w.w.Write([]byte(x))
}

// TODO: Identifier

func (w *Writer) VarInt(x int32) {
	for x >= 0x80 {
		_ = w.w.WriteByte(byte(x&0x7f) | 0x80)
		x >>= 7
	}

	_ = w.w.WriteByte(byte(x))
}

func (w *Writer) VarLong(x int64) {
	for x >= 0x80 {
		_ = w.w.WriteByte(byte(x&0x7f) | 0x80)
		x >>= 7
	}

	_ = w.w.WriteByte(byte(x))
}
