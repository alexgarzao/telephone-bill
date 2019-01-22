package domain

import (
	"strconv"
	"time"
)

type StartRecord struct {
	recordID    string
	timestamp   time.Time
	callID      string
	source      string
	destination string
}

func NewStartRecord(recordID string, timestamp time.Time, callID string, source string, destination string) *StartRecord {
	if (recordID == "" || timestamp == time.Time{} || callID == "" || source == "" || destination == "") {
		return nil
	}

	if !isValidNumber(source) || !isValidNumber(destination) {
		return nil
	}

	return &StartRecord{
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
