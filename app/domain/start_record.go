package domain

import "time"

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

	return &StartRecord{
		recordID:    recordID,
		timestamp:   timestamp,
		callID:      callID,
		source:      source,
		destination: destination,
	}
}
