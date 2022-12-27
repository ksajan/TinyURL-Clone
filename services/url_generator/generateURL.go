package service

import (
	"fmt"
	"strings"
	utils "tinyURL/utils"
)

// GenerateURL generates a URL based on the given parameters

func GenerateURL(LongURL string, baseDomainUrl string) (string, string) {
	fmt.Println("Generating URL")
	bEncodedValue := utils.Base64URLEncoding(utils.GenerateUUID(LongURL))
	return strings.Join([]string{baseDomainUrl, bEncodedValue[:7]}, "/"), bEncodedValue[:7]
}
