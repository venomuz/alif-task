package hash

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	String(str string) (string, error)
	CheckString(hashedPassword, password string) error
}

type Hash struct{}

func NewPasswordHasher() *Hash {
	return &Hash{}
}

func (h *Hash) String(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (h *Hash) CheckString(hashedPassword, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}
