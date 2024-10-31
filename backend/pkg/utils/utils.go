package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func GenerateUniqueID() (string, error) {
	b := make([]byte, 16)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	uniqueID := base64.RawURLEncoding.EncodeToString(b)
	uniqueID = strings.ReplaceAll(uniqueID, "=", "")
	uniqueID = strings.ReplaceAll(uniqueID, "_", "")
	uniqueID = strings.ReplaceAll(uniqueID, "-", "")
	uniqueID = strings.ToLower(uniqueID)

	return uniqueID[:12], nil
}
