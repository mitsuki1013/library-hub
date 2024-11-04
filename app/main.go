package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"library-hub/app/adapter/api/action"
	"library-hub/app/adapter/repository"
	"library-hub/app/infrastructure/database"
	"library-hub/app/usecase"
	"log"
	"net/http"
	"time"
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

	r := mux.NewRouter()
	uc := usecase.NewCreateCompanyInteractor(
		repository.NewCompanyMySQL(mysql.Db),
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
