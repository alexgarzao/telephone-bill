package domain

import (
	"time"
)

type StopCall struct {
	recordID  string
	timestamp time.Time
	callID    string
}

func NewStopCall(recordID string, timestamp time.Time, callID string) *StopCall {
	if (recordID == "" || timestamp == time.Time{} || callID == "") {
		return nil
	}

	return &StopCall{
		recordID:  recordID,
		timestamp: timestamp,
		callID:    callID,
	}
}
