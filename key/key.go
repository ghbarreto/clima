package key

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Env() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env file couldn't be loaded")
	}

	return os.Getenv("API_KEY")
}
