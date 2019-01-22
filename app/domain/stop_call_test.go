package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStopCall(t *testing.T) {
	n := time.Now()

	tables := []StopCall{
		{"R1", n, "C1"},
		{"R1123", n, "asasdasd"},
		{"999", n, "23123132@#@!@@"},
		{"aaAAA321312", n, "C1"},
	}

	for _, table := range tables {
		r := NewStopCall(table.recordID, table.timestamp, table.callID)
		assert.NotNil(t, r, "Must be a valid object!")
		assert.Equal(t, r.recordID, table.recordID)
		assert.Equal(t, r.timestamp, table.timestamp)
		assert.Equal(t, r.callID, table.callID)
	}
}

func TestInvalidStopCallWhenSomeFieldIsEmpty(t *testing.T) {
	n := time.Now()

	tables := []StopCall{
		{"", n, "C1"},
		{"R1", time.Time{}, "C1"},
		{"R1", n, ""},
	}

	for _, table := range tables {
		r := NewStopCall(table.recordID, table.timestamp, table.callID)
		assert.Nil(t, r, "Must be a invalid object!")
	}
}
