package main

import (
	"log"
	"shopping-cart/inits"
	"shopping-cart/models"
)

func init() {
	inits.LoadEnvVariables()
	inits.ConnectToDB()
}

func main() {
	err := inits.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Cart{},
		&models.Order{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful")
	}
}