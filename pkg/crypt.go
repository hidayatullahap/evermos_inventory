package pkg

import (
	"crypto/md5"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}

// GenerateToken returns a unique token based on the provided email string
func GenerateToken(username string) (string, error) {
	hash, err := HashAndSalt([]byte(username))
	if err != nil {
		return "", err
	}

	md := md5.New()
	md.Write(hash)

	return hex.EncodeToString(md.Sum(nil)), nil
}
