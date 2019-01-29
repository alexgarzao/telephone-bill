package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StartCall struct {
	RecordID    string
	Timestamp   time.Time
	CallID      string
	Source      string
	Destination string
}

type RecordStartCallInteractor interface {
	Add(recordID string, timestamp time.Time, callID string, source string, destination string) error
}

type RestAPIHandler struct {
	RecordStartCallInteractor RecordStartCallInteractor
}

func (handler RestAPIHandler) Add(res http.ResponseWriter, req *http.Request) {
	var startCall StartCall
	if err := json.NewDecoder(req.Body).Decode(&startCall); err != nil {
		respondWithError(res, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %s", err.Error()))
		return
	}
	err := handler.RecordStartCallInteractor.Add(startCall.RecordID, startCall.Timestamp, startCall.CallID, startCall.Source, startCall.Destination)
	if err != nil {
		respondWithError(res, http.StatusBadRequest, fmt.Sprintf("Invalid data: %s", err.Error()))
		return
	}
	json.NewEncoder(res).Encode(startCall)
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
