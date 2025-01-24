package functions

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) (*string, error) {

	decodedString, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, err
	}
	encodedToBase64 := base64.StdEncoding.EncodeToString(decodedString)
	return &encodedToBase64, nil
}
