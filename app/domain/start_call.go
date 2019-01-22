package domain

import (
	"strconv"
	"time"
)

type StartCall struct {
	recordID    string
	timestamp   time.Time
	callID      string
	source      string
	destination string
}

func NewStartCall(recordID string, timestamp time.Time, callID string, source string, destination string) *StartCall {
	if (recordID == "" || timestamp == time.Time{} || callID == "" || source == "" || destination == "") {
		return nil
	}

	if !isValidNumber(source) || !isValidNumber(destination) {
		return nil
	}

	return &StartCall{
		recordID:    recordID,
		timestamp:   timestamp,
		callID:      callID,
		source:      source,
		destination: destination,
	}
}

func isValidNumber(number string) bool {
	_, err := strconv.ParseUint(number, 10, 64)
	return (err == nil && (len(number) == 10 || len(number) == 11))
}
