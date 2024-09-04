package ctxt

import (
	"context"
	"testing"

	"github.com/billykore/kore/backend/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestContextWithUserAndUserFromContext(t *testing.T) {
	ctx := context.Background()
	uctx := ContextWithUser(ctx, entity.User{
		Username: "user",
	})
	assert.NotNil(t, uctx)

	user, ok := UserFromContext(uctx)
	assert.True(t, ok)
	assert.Equal(t, "user", user.Username)
}
