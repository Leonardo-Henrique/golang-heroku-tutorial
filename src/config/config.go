package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT = "7000"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	if p := os.Getenv("PORT"); p != "" {
		PORT = p
	}

}
