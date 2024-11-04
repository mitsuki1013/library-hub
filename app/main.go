package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"library-hub/app/adapter/api/action"
	"library-hub/app/adapter/repository"
	"library-hub/app/usecase"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	db := &gorm.DB{}
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
