package rabbit

import (
	"testing"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestMarshalBinary(t *testing.T) {
	payload := &Payload[*entity.UpdateShippingRabbitData]{
		Origin: "test-origin",
		Data: &entity.UpdateShippingRabbitData{
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
	payload := new(Payload[*entity.UpdateShippingRabbitData])
	err := payload.UnmarshalBinary([]byte(s))
	assert.NoError(t, err)
	assert.Equal(t, "test-origin", payload.Origin)
}
