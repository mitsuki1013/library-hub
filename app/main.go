package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"library-hub/app/adapter/api/action"
	"library-hub/app/adapter/repository"
	"library-hub/app/usecase"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PW"),
		os.Getenv("DB_HOST_PORT"),
		os.Getenv("MYSQL_DB"))
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	r := mux.NewRouter()
	uc := usecase.NewCreateCompanyInteractor(
		repository.NewCompanyMySQL(db),
	)
	act := action.NewCreateCompanyAction(uc)
	r.HandleFunc("/company", act.CreateCompany).Methods(http.MethodPost)

	// サーバ設定
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
