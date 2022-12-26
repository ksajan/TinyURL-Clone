package utils

import (
	b64 "encoding/base64"
	"fmt"

	uuid "github.com/google/uuid"
)

// GenerateUUID generates a UUID

func GenerateUUID(data string) string {
	fmt.Println("Generating UUID version 3 using MD5")
	uuidBytes := uuid.NewMD5(uuid.NameSpaceURL, []byte(data))
	fmt.Printf("UUID: %s", uuidBytes.String())
	return uuidBytes.String()
}

func Base64URLEncoding(data string) string {
	fmt.Println("Generating Base64 Encoding")
	sEncoded := b64.URLEncoding.EncodeToString([]byte(data))
	return sEncoded
}

func Base64URLDecoding(data string) string {
	fmt.Println("Generating Base64 Decoding")
	sDecoded, _ := b64.URLEncoding.DecodeString(data)
	return string(sDecoded)
}
