package service

import (
	"fmt"
	"strings"
	utils "tinyURL/utils"
)

// GenerateURL generates a URL based on the given parameters

func GenerateURL(url string, basePath string) string {
	fmt.Println("Generating URL")
	bEncodedValue := utils.Base64URLEncoding(utils.GenerateUUID(url))
	return strings.Join([]string{basePath, bEncodedValue[:7]}, "/")
}
