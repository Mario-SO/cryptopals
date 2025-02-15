package convert_hex_to_base_64

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a hexadecimal string to base64
func HexToBase64(hexStr string) (string, error) {
	// Decode hex string to bytes
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	// Encode bytes to base64
	return base64.StdEncoding.EncodeToString(bytes), nil
}
