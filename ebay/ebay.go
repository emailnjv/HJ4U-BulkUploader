package ebay

import (
	"os"

	"github.com/joho/godotenv"
)

type EbayClient struct {
	oAuthKey string
}

// NewEbayClient instantiates, and returns a new Ebay client
func NewEbayClient() (EbayClient, error) {
	var result EbayClient
	err := godotenv.Load("../.env")
	if err != nil {
		return result, err
	}
	result.oAuthKey = os.Getenv("EBAY_OAUTH_TOKEN")
	if result.oAuthKey != "" {
		return result, err
	}

	return result, err
}


