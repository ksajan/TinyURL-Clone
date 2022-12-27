package api

import (
	"fmt"
	os "os"
	time "time"
	db "tinyURL/Database/Airospike"
	tinyURL "tinyURL/services/url_generator"

	aero "github.com/aerospike/aerospike-client-go"
	fiber "github.com/gofiber/fiber/v2"
)

func ShortenURL(c *fiber.Ctx) error {
	body := new(ShortURLRequest)
	body.CreatedAt = time.Now()
	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	// Validate the URL
	ValidateURL(body.LongURL, c)

	/*
		1. Check if the URL is already shortened
		2. If yes, return the shortened URL with addition to db
		3. If no, generate a new URL and return it with addition to db
	*/
	// baseURL := []string{os.Getenv("BASE_URL"), os.Getenv("APP_PORT")}
	ShortURL, bEncodedValue := tinyURL.GenerateURL(body.LongURL, os.Getenv("BASE_URL"))
	// bEncodedValue := utils.Base64URLEncoding(utils.GenerateUUID(body.LongURL))

	if !db.CheckRecordExists(bEncodedValue) {
		shortURLBinMap := aero.BinMap{
			"longURL":    body.LongURL,
			"shortURL":   ShortURL,
			"created_at": string(body.CreatedAt.Format("2006-01-02 15:04:05")),
			"updated_at": string(time.Now().Format("2006-01-02 15:04:05")),
			"expires_at": string(time.Now().Add(time.Hour * time.Duration(24*60)).Format("2006-01-02 15:04:05")),
		}
		// Add to DB and return the short URL
		db.AddToRecord(bEncodedValue, shortURLBinMap)
		return c.JSON(fiber.Map{
			"shortURL": ShortURL,
		})
	} else {
		// Get the short URL from DB and return it
		record, err := db.ReadRecord(bEncodedValue)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "cannot read record",
			})
		}
		body.ShortURL = record.Bins["shortURL"].(string)
		return c.JSON(fiber.Map{
			"shortURL": body.ShortURL,
		})
	}
}

func ResolveURL(c *fiber.Ctx) error {
	// body := new(ShortURLRequest)
	// Get the short URL from the request
	shortURL := c.Params("shortURLEncodedValue")
	fmt.Println("shortURL: ", shortURL)
	// Get the long URL from the DB
	record, err := db.ReadRecord(shortURL)
	fmt.Println("record: ", record)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot read record",
		})
	}
	return c.Redirect(record.Bins["longURL"].(string), fiber.StatusTemporaryRedirect)
}
