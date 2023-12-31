package password

import "golang.org/x/crypto/bcrypt"

func Verify(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
