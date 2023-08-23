package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword is func to encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash is func to validate password hasher
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsNumber is func to validate type from input string is number or alphabet
func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func GenerateHexadecimalStringTokent() (*string, error) {
	hashLength := 20 // 20 bytes = 40 hexadecimal characters

	// Generate random bytes
	randomBytes := make([]byte, hashLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Convert random bytes to hexadecimal string
	hashString := hex.EncodeToString(randomBytes)
	return &hashString, nil
}

func ParseTokenHex(token string) (string, error) {
	substrings := strings.Split(token, " ")
	trimmedStr := strings.TrimSpace(substrings[1])

	return trimmedStr, nil
}

func GetCustomerXidFromToken(rdsKey string) string {
	substrings := strings.Split(rdsKey, " ")
	trimmedStr := strings.TrimSpace(substrings[1])

	return trimmedStr
}
