package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

func HashUrl(inputURL string) string {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		panic(err)
	}

	const queryParams, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Specktor Parse Query Parameters:")
	fmt.Println(queryParams)

	hash := sha256.New()
	hash.Write([]byte(inputURL))
	specterURL := hash.Sum(nil)

	return hex.EncodeToString(specterURL)
}

func main() {
	inputURL := ""
	hashedURL := HashUrl(inputURL)
	fmt.Println("FafaHA:", hashedURL)
}
