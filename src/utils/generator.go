package utils

import (
	"encoding/json"

	"github.com/Luzifer/go-openssl/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	SECRET_KEY = "secret"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func EncryptDevice(device map[string]interface{}) (string, error) {
	credentials := openssl.BytesToKeyMD5
	jsonDevice, err := json.Marshal(device)
	if err != nil {
		return "", err
	}
	byteDevice, err := openssl.New().EncryptBytes(SECRET_KEY, jsonDevice, credentials)
	if err != nil {
		return "", err
	}
	return string(byteDevice), nil
}
