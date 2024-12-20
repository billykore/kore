package otp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	otp := Generate(6)
	assert.Len(t, otp.Value, 6)

	otp2 := Generate(10)
	assert.Len(t, otp2.Value, 10)
}

func BenchmarkGenerate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Generate(6)
	}
}
