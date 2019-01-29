package domain

import (
	"fmt"
	"time"
)

type StopCall struct {
	recordID  string
	timestamp time.Time
	callID    string
}

func NewStopCall(recordID string, timestamp time.Time, callID string) (*StopCall, error) {
	if (recordID == "" || timestamp == time.Time{} || callID == "") {
		return nil, fmt.Errorf("empty fields")
	}

	return &StopCall{
		recordID:  recordID,
		timestamp: timestamp,
		callID:    callID,
	}, nil
}
