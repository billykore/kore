package rabbit

import "encoding/json"

// Payload is message broker payload.
type Payload[T any] struct {
	// Origin is where the payload come from.
	Origin string `json:"origin"`
	// Data to publish or consume from the message broker.
	Data T `json:"data"`
}

// MarshalJSON return JSON encoding of Payload.
func (p *Payload[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}

// UnmarshalJSON parses encoded data into Payload.
func (p *Payload[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, p)
}
