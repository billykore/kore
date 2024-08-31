package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationErr(t *testing.T) {
	type s struct {
		F string `validate:"required"`
	}
	empty := s{}
	v := New()
	err := v.Validate(empty)
	assert.Error(t, err)
	assert.Equal(t, "F is required", err.Error())
}
