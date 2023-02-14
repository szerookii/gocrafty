package types

import (
	"encoding/json"
	"log"
)

type Chat struct {
	Text          string `json:"text,omitempty"`
	Bold          bool   `json:"bold,omitempty"`
	Italic        bool   `json:"italic,omitempty"`
	Underlined    bool   `json:"underlined,omitempty"`
	Strikethrough bool   `json:"strikethrough,omitempty"`
	Obfuscated    bool   `json:"obfuscated,omitempty"`
	Color         string `json:"color,omitempty"`
	Translate     string `json:"translate,omitempty"`
	With          []Chat `json:"with,omitempty"`
	Extra         []Chat `json:"extra,omitempty"`
}

func (m Chat) JSON() []byte {
	code, err := json.Marshal(m)
	if err != nil {
		log.Panicln(err)
	}

	return code
}
