package helpers

import "golang.org/x/crypto/bcrypt"

func HassPass(password string) string {
	salt := 8
	p := []byte(password)
	hash, _ := bcrypt.GenerateFromPassword(p, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
