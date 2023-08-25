package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

func TryHw3(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		panic(err)
	}

	queryParams, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Spector Parsi Query Parameters:")
	fmt.Println(queryParams)

	hash := sha256.New()
	hash.Write([]byte(inputURL))
	specterURL := hash.Sum(nil)

	return hex.EncodeToString(specterURL)
}

func main() {
	inputURL := ""
	hashedURL := TryHw3(inputURL)
	fmt.Println("FafaHA:", hashedURL)
}