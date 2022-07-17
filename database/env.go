package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ENVLoad() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("ENV NOT OPEN")
		fmt.Println("ENV NOT OPEN")
	}

	return os.Getenv("MONGOURI")
}
