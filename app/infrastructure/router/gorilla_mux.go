package router

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

func Listen(db *gorm.DB) {
	srv := &http.Server{
		Handler:      setAppHandlers(db),
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func setAppHandlers(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	act := action.NewCreateCompanyAction(
		usecase.NewCreateCompanyInteractor(
			repository.NewCompanyMySQL(db),
		))
	r.HandleFunc("/company", act.CreateCompany).Methods(http.MethodPost)

	return r
}
