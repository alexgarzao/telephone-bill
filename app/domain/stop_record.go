package domain

import (
	"time"
)

type StopRecord struct {
	recordID  string
	timestamp time.Time
	callID    string
}

func NewStopRecord(recordID string, timestamp time.Time, callID string) *StopRecord {
	if (recordID == "" || timestamp == time.Time{} || callID == "") {
		return nil
	}

	return &StopRecord{
		recordID:  recordID,
		timestamp: timestamp,
		callID:    callID,
	}
}
