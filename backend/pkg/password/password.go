package password

import "golang.org/x/crypto/bcrypt"

// Verify compare hashed and string password.
func Verify(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
