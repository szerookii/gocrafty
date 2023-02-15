package login

import (
	"github.com/szerookii/gocrafty/gocrafty/minecraft/protocol"
)

type EncryptionRequest struct {
	ServerID          string
	PublicKeyLength   int32
	PublicKey         []byte
	VerifyTokenLength int32
	VerifyToken       []byte
}

func (d *EncryptionRequest) ID() int32 {
	return IDEncryptionRequest
}

func (s *EncryptionRequest) State() int32 {
	return 0 // not used
}

func (d *EncryptionRequest) Marshal(w *protocol.Writer) {
	w.String(d.ServerID)
	w.VarInt(d.PublicKeyLength)
	w.WriteBytes(d.PublicKey)
	w.VarInt(d.VerifyTokenLength)
	w.WriteBytes(d.VerifyToken)
}

func (d *EncryptionRequest) Unmarshal(r *protocol.Reader) {}
