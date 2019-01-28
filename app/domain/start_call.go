package domain

import (
	"strconv"
	"time"
)

type StartCallRepository interface {
	Store(startCall StartCall)
	FindById(id int) StartCall
}

type StartCall struct {
	RecordID    string
	Timestamp   time.Time
	CallID      string
	Source      string
	Destination string
}

func NewStartCall(recordID string, timestamp time.Time, callID string, source string, destination string) *StartCall {
	if (recordID == "" || timestamp == time.Time{} || callID == "" || source == "" || destination == "") {
		return nil
	}

	if !isValidNumber(source) || !isValidNumber(destination) {
		return nil
	}

	return &StartCall{
		RecordID:    recordID,
		Timestamp:   timestamp,
		CallID:      callID,
		Source:      source,
		Destination: destination,
	}
}

func isValidNumber(number string) bool {
	_, err := strconv.ParseUint(number, 10, 64)
	return (err == nil && (len(number) == 10 || len(number) == 11))
}
