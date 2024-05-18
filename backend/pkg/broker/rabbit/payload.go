package rabbit

import "encoding/json"

type Payload[T any] struct {
	Origin string `json:"origin"`
	Data   T      `json:"data"`
}

func (p *Payload[T]) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Payload[T]) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, p)
}
