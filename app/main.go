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

	recordCallInteractor := new(usecases.RecordCallInteractor)
	recordCallInteractor.Logger = new(infrastructure.Logger)
	recordCallInteractor.StartCallRepository, err = interfaces.NewDbStartCallRepo(dbHandler)
	recordCallInteractor.StopCallRepository, err = interfaces.NewDbStopCallRepo(dbHandler)
	if err != nil {
		log.Fatalf("Error when trying to create DbStartCallRepo: %s\n", err.Error())
	}

	restAPIHandler := interfaces.RestAPIHandler{}
	restAPIHandler.RecordCallInteractor = recordCallInteractor

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/billing/recordcalls", func(res http.ResponseWriter, req *http.Request) {
		restAPIHandler.RecordCall(res, req)
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
