package inits

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	fmt.Println("Envs Loaded")
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}