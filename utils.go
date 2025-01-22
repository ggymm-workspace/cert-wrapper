package main

import (
	"encoding/base64"
)

func encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func decode(input string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(input)
	return decoded
}
