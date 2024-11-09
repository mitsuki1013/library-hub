package main

import (
	"github.com/joho/godotenv"
	"library-hub/app/infrastructure/database"
	"library-hub/app/infrastructure/router"
	"log"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mysql, err := database.NewMysqlHandler(database.NewConfigMySQL())
	if err != nil {
		log.Fatalf("Error connecting to mysql: %v", err)
	}

	router.Listen(mysql.Db)
}
