package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStartRecord(t *testing.T) {
	n := time.Now()

	tables := []StartRecord{
		{"R1", n, "C1", "S1", "D1"},
	}

	for _, table := range tables {
		r := NewStartRecord(table.recordID, table.timestamp, table.callID, table.source, table.destination)
		assert.NotNil(t, r, "Must be a valid object!")
		assert.Equal(t, r.recordID, table.recordID)
		assert.Equal(t, r.timestamp, table.timestamp)
		assert.Equal(t, r.callID, table.callID)
		assert.Equal(t, r.source, table.source)
		assert.Equal(t, r.destination, table.destination)
	}
}

func TestInvalidWhenSomeFieldIsEmpty(t *testing.T) {
	n := time.Now()

	tables := []StartRecord{
		{"", n, "C1", "S1", "D1"},
		{"R1", time.Time{}, "C1", "S1", "D1"},
		{"R1", n, "", "S1", "D1"},
		{"R1", n, "C1", "", "D1"},
		{"R1", n, "C1", "S1", ""},
	}

	for _, table := range tables {
		r := NewStartRecord(table.recordID, table.timestamp, table.callID, table.source, table.destination)
		assert.Nil(t, r, "Must be a invalid object!")
	}
}
