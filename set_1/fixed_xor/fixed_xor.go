package fixedxor

import (
	"encoding/hex"
)

func FixedXOR(s string) (string, error) {
	against := "686974207468652062756c6c277320657965"
	sDecoded, err := hex.DecodeString(s)
	if err != nil {
		return "error decoding to hex variable against", err
	}

	againstDecoded, err := hex.DecodeString(against)
	if err != nil {
		return "error decoding to hex variable against", err
	}

	result := make([]byte, len(sDecoded))

	for i := 0; i < len(sDecoded); i++ {
		result[i] = sDecoded[i] ^ againstDecoded[i]
	}

	return hex.EncodeToString(result), nil
}
