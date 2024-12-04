package feature

import (
	"errors"
	"os"
	"strconv"
)

var ErrFeatureIsDisabled = errors.New("feature is disabled")

func IsEnabled(flag string) bool {
	value, exists := os.LookupEnv(flag)
	if !exists {
		return false
	}
	enabled, err := strconv.ParseBool(value)
	return err == nil && enabled
}
