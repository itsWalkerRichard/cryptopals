package functions

import (
	"encoding/hex"
	"fmt"
)

func FixedXOR(hex1, hex2 string) (*string, error) {

	decodedString1, err := hex.DecodeString(hex1)
	if err != nil {
		return nil, err
	}
	decodedString2, err := hex.DecodeString(hex2)
	if err != nil {
		return nil, err
	}
	if len(decodedString1) != len(decodedString2) {
		return nil, fmt.Errorf("length of hex1 and hex2 is not equal")
	}

	for i := range decodedString1 {
		decodedString1[i] = decodedString1[i] ^ decodedString2[i]
	}

	xorString := hex.EncodeToString(decodedString1)

	return &xorString, nil
}
