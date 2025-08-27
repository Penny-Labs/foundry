package hash

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

const (
	BcryptCost = 12 // Cost for bcrypt hashing
)

// Sha512Hash hashes the provided key using SHA-512 and returns the hex representation.
func Sha512Hash(key string) string {
	hash := sha512.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// Sha256Hash hashes the provided key using SHA-256 and returns the hex representation.
func Sha256Hash(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	return hex.EncodeToString(hash.Sum(nil))
}

// BcryptHashKey hashes the provided key using bcrypt after hashing it with SHA-256.
func BcryptHashKey(key string, cost int) (string, error) {
	sha := Sha256Hash(key)
	shaBytes := []byte(sha)
	hashed, err := bcrypt.GenerateFromPassword(shaBytes, cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// BcryptCompareHashAndKey compares a bcrypt hashed key with a plain key after hashing it with SHA-256.
func BcryptCompareHashAndKey(hashedKey, key string) error {
	sha := Sha256Hash(key)
	return bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(sha))
}
