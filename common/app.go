package common

import (
	"log"
	"github.com/joho/godotenv"
	metalabs "github.com/meta-labs/meta-labs-go/metalabs_sdk"
)

var (
	// Client (Linter)
	Client metalabs.Client
)

// LoadEnv (Linter)
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env Not Found")
	}
}