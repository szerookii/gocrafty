package protocol

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/szerookii/gocrafty/gocrafty/minecraft/types"
)

type Writer struct {
	bytes.Buffer
}

func (w *Writer) Bytes() []byte {
	return w.Buffer.Bytes()[:w.Buffer.Len()]
}

func (w *Writer) WriteBytes(data []byte) {
	w.Buffer.Write(data)
}

func (w *Writer) Bool(x bool) {
	if x {
		_ = w.WriteByte(0x01)
	} else {
		_ = w.WriteByte(0x00)
	}
}

func (w *Writer) Short(x int16) {
	w.Grow(2)
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) UShort(x uint16) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) Int(x int32) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) UInt(x uint32) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) Long(x int64) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) ULong(x uint64) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) Float(x float32) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) Double(x float64) {
	_ = binary.Write(&w.Buffer, binary.BigEndian, x)
}

func (w *Writer) String(x string) {
	w.VarInt(int32(len(x)))
	w.Buffer.WriteString(x)
}

// TODO: Identifier

func (w *Writer) VarInt(x int32) {
	raw := uint32(x)

	for {
		temp := byte(raw & 0b01111111)
		raw >>= 7

		if raw != 0 {
			temp |= 0b10000000
		}

		_ = w.WriteByte(temp)

		if raw == 0 {
			break
		}
	}
}

func (w *Writer) VarLong(x int64) {
	raw := uint64(x)

	for {
		temp := byte(raw & 0b01111111)
		raw >>= 7

		if raw != 0 {
			temp |= 0b10000000
		}

		w.WriteByte(temp)

		if raw == 0 {
			break
		}
	}
}

func (w *Writer) WriteUUID(val uuid.UUID) {
	b, _ := val.MarshalBinary()
	w.WriteBytes(b)
}

func (w *Writer) Chat(x types.Chat) {
	b := x.JSON()

	w.VarInt(int32(len(b)))
	w.WriteBytes(b)
}

func (w *Writer) JSON(x any) {
	b, err := json.Marshal(x)

	if err != nil {
		panic(err)
	}

	w.VarInt(int32(len(b)))
	w.WriteBytes(b)
}
