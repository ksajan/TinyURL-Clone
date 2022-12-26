package aerospike

import time "time"

type Form struct {
	LongURL string `json:"long_url"`
}

type TinyURL struct {
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	Expiry    time.Time `json:"expiry"`
	CreatedAt time.Time `json:"created_at"`
}

type TinyURLResponse struct {
	ShortURL string `json:"short_url"`
}

type TinyURLVisits struct {
	ShortURL string `json:"short_url"`
	Visits   int    `json:"visits"`
}

type TinyURLRequestCount struct {
	ShortURL string `json:"short_url"`
	Count    int    `json:"count"`
	Date     string `json:"date"`
}

type ErrorMessages struct {
	ErrorMessage string `json:"error_message"`
}
