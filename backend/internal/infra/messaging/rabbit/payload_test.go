package rabbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	ShippingId uint   `json:"shippingId"`
	Status     string `json:"status"`
}

func TestMarshalBinary(t *testing.T) {
	payload := &Payload[*testData]{
		Origin: "test-origin",
		Data: &testData{
			ShippingId: 69,
			Status:     "shipped",
		},
	}
	b, err := payload.MarshalBinary()
	assert.NoError(t, err)
	assert.Equal(t, `{"origin":"test-origin","data":{"shippingId":69,"status":"shipped"}}`, string(b))
}

func TestUnmarshalBinary(t *testing.T) {
	s := `{"origin":"test-origin","data":{"shippingId":69,"status":"shipped"}}`
	payload := new(Payload[*testData])
	err := payload.UnmarshalBinary([]byte(s))
	assert.NoError(t, err)
	assert.Equal(t, "test-origin", payload.Origin)
}
