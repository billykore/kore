package rabbit

import "encoding/json"

type Payload struct {
	Origin string `json:"origin"`
	Data   any    `json:"data"`
}

func NewPayload(origin string, data any) *Payload {
	return &Payload{
		Origin: origin,
		Data:   data,
	}
}

func (p *Payload) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}
