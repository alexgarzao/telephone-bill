package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Call struct {
	RecordID    string
	Type        string
	Timestamp   time.Time
	CallID      string
	Source      string
	Destination string
}

type RecordCallInteractor interface {
	AddStart(recordID string, timestamp time.Time, callID string, source string, destination string) error
	AddStop(recordID string, timestamp time.Time, callID string) error
}

type RestAPIHandler struct {
	RecordCallInteractor RecordCallInteractor
}

func (handler RestAPIHandler) RecordCall(res http.ResponseWriter, req *http.Request) {
	var err error
	var call Call
	if err := json.NewDecoder(req.Body).Decode(&call); err != nil {
		respondWithError(res, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}

	if call.Type == "start" {
		err = handler.RecordCallInteractor.AddStart(call.RecordID, call.Timestamp, call.CallID, call.Source, call.Destination)
	} else {
		err = handler.RecordCallInteractor.AddStop(call.RecordID, call.Timestamp, call.CallID)
	}
	if err != nil {
		respondWithError(res, http.StatusBadRequest, fmt.Sprintf("Invalid data: %s", err.Error()))
		return
	}
	json.NewEncoder(res).Encode(call)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
