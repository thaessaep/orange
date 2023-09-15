package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Token string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("cannot load token")
	}

	Token = os.Getenv("TOKEN")
}
