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
	return &StartRecord{}
}
