package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/billykore/kore/backend/domain/user"
	"github.com/redis/go-redis/v9"
)

const failedLoginKey = "login_failed"

type UserCache struct {
	redis *redis.Client
}

func NewUserCache(redis *redis.Client) *UserCache {
	return &UserCache{
		redis: redis,
	}
}

func (c *UserCache) SaveFailedLoginAttempt(ctx context.Context, authActivity user.AuthActivities) error {
	b, err := json.Marshal(authActivity)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("%s:%s", failedLoginKey, authActivity.Username)
	cmd := c.redis.Set(ctx, key, b, 24*time.Hour)
	return cmd.Err()
}

func (c *UserCache) GetFailedLoginAttempt(ctx context.Context, username string) (*user.AuthActivities, error) {
	key := fmt.Sprintf("%s:%s", failedLoginKey, username)
	val, err := c.redis.Get(ctx, key).Bytes()
	if err != nil && errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	activity := new(user.AuthActivities)
	err = activity.UnmarshalBinary(val)
	if err != nil {
		return nil, err
	}
	return activity, nil
}
