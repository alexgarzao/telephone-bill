package domain

import (
	"fmt"
	"time"
)

type StopCall struct {
	RecordID  string
	Timestamp time.Time
	CallID    string
}

func NewStopCall(recordID string, timestamp time.Time, callID string) (*StopCall, error) {
	if (recordID == "" || timestamp == time.Time{} || callID == "") {
		return nil, fmt.Errorf("empty fields")
	}

	return &StopCall{
		RecordID:  recordID,
		Timestamp: timestamp,
		CallID:    callID,
	}, nil
}
