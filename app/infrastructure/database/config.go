package database

import "os"

type config struct {
	host     string
	database string
	port     string
	user     string
	password string
}

func NewConfigMySQL() *config {
	return &config{
		host:     "localhost",
		database: os.Getenv("MYSQL_DB"),
		port:     os.Getenv("DB_HOST_PORT"),
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PW"),
	}
}
