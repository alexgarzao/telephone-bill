package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexgarzao/telephone-bill/app/interfaces"
	"github.com/alexgarzao/telephone-bill/app/usecases"
	"github.com/gorilla/mux"

	"github.com/alexgarzao/telephone-bill/app/infrastructure"
)

func main() {
	fmt.Println("Starting!")

	dbHandler, err := infrastructure.NewSqlite("/var/tmp/production.sqlite")
	if err != nil {
		log.Fatalf("Error when trying to connect to database: %s\n", err.Error())
	}

	recordStartCallInteractor := new(usecases.RecordStartCallInteractor)
	recordStartCallInteractor.StartCallRepository, err = interfaces.NewDbStartCallRepo(dbHandler)
	recordStartCallInteractor.Logger = new(infrastructure.Logger)

	if err != nil {
		log.Fatalf("Error when trying to create DbStartCallRepo: %s\n", err.Error())
	}

	restAPIHandler := interfaces.RestAPIHandler{}
	restAPIHandler.RecordStartCallInteractor = recordStartCallInteractor

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/billing/startcalls", func(res http.ResponseWriter, req *http.Request) {
		restAPIHandler.Add(res, req)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
