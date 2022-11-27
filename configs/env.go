package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load("configs/.dev.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	return os.Getenv("MONGO_URI")
}
