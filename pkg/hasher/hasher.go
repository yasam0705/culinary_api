package hasher

import "golang.org/x/crypto/bcrypt"

const defaultCost = 9

func Generate(s string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(s), defaultCost)
	return string(pass), err
}

func Compare(pass1, pass2 string) error {
	return bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2))
}
