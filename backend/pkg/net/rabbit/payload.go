package rabbit

import "encoding/json"

// Payload is message broker payload.
type Payload[T any] struct {
	// Origin is where the payload come from.
	Origin string `json:"origin"`
	// Data to publish or consume from the message broker.
	Data T `json:"data"`
}

// MarshalBinary return JSON encoding of Payload.
func (p *Payload[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalBinary parses encoded data into Payload.
func (p *Payload[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
