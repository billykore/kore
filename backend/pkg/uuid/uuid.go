package uuid

import "github.com/google/uuid"

// New return new string UUID version 7.
func New() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
