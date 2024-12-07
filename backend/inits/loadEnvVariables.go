package inits

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	fmt.Println("Envs Loading")
	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Fatal("Error: .env file does not exist in the working directory")
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}