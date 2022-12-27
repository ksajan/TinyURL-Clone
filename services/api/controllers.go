package api

import (
	urlParser "net/url"

	validators "github.com/asaskevich/govalidator"
	fiber "github.com/gofiber/fiber/v2"
)

func ValidateURL(url string, c *fiber.Ctx) (bool, error) {
	isValid := validators.IsURL(url)
	if !isValid {
		return false, c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}
	return true, nil
}

func GetEncodedValuefromURL(url string, c *fiber.Ctx) (string, error) {
	parsedURL, err := urlParser.Parse(url)
	if err != nil {
		return "", c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}
	return parsedURL.Path, nil
}
