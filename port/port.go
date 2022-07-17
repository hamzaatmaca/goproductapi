package port

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func PORTLoad() string {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("ENV NOT OPEN")
		fmt.Println("ENV NOT OPEN")
	}

	return os.Getenv("PORT")
}
