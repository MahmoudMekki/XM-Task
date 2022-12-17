package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return ""
	}
	return string(hashed)
}
func ComparePasswordHashed(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash, plain)
	if err != nil {
		return false
	}
	return true
}
