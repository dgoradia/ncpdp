package ncpdp

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
)

func NewMessage(message string) (*Message, error) {
	dec, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return nil, err
	}

	var msg Message
	return &msg, xml.Unmarshal(dec, &msg)
}

func (m *Message) ToJson() ([]byte, error) {
	return json.Marshal(m)
}
