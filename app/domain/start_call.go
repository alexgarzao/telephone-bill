package domain

import (
	"fmt"
	"strconv"
	"time"
)

type StartCallRepository interface {
	Store(startCall StartCall) error
	FindById(id int) (*StartCall, error)
}

type StartCall struct {
	RecordID    string
	Timestamp   time.Time
	CallID      string
	Source      string
	Destination string
}

func NewStartCall(recordID string, timestamp time.Time, callID string, source string, destination string) (*StartCall, error) {
	if (recordID == "" || timestamp == time.Time{} || callID == "" || source == "" || destination == "") {
		return nil, fmt.Errorf("empty fields")
	}

	if !isValidNumber(source) {
		return nil, fmt.Errorf("invalid source number: %s", source)
	}

	if !isValidNumber(destination) {
		return nil, fmt.Errorf("invalid destination number: %s", destination)
	}

	return &StartCall{
		RecordID:    recordID,
		Timestamp:   timestamp,
		CallID:      callID,
		Source:      source,
		Destination: destination,
	}, nil
}

func isValidNumber(number string) bool {
	_, err := strconv.ParseUint(number, 10, 64)
	return (err == nil && (len(number) == 10 || len(number) == 11))
}
