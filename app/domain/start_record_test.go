package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStartRecord(t *testing.T) {
	n := time.Now()

	tables := []StartRecord{
		{"R1", n, "C1", "1212345678", "12123456789"},
		{"R1123", n, "asasdasd", "1232345678", "12123456789"},
		{"999", n, "23123132@#@!@@", "1212222678", "12123456789"},
		{"aaAAA321312", n, "C1", "12123456222", "12123456789"},
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

func TestInvalidStartRecordWhenSomeFieldIsEmpty(t *testing.T) {
	n := time.Now()

	tables := []StartRecord{
		{"", n, "C1", "1212345678", "12123456789"},
		{"R1", time.Time{}, "C1", "1212345678", "12123456789"},
		{"R1", n, "", "1212345678", "12123456789"},
		{"R1", n, "C1", "", "12123456789"},
		{"R1", n, "C1", "1212345678", ""},
	}

	for _, table := range tables {
		r := NewStartRecord(table.recordID, table.timestamp, table.callID, table.source, table.destination)
		assert.Nil(t, r, "Must be a invalid object!")
	}
}

func TestInvalidStartRecordWhenSomeFieldIsInvalid(t *testing.T) {
	n := time.Now()

	tables := []StartRecord{
		{"R1", n, "C1", "12145678", "12123456789"},
		{"R1", n, "C1", "1212345678", "1212789"},
		{"R1", n, "C1", "1212345678", "1212789123123"},
	}

	for _, table := range tables {
		r := NewStartRecord(table.recordID, table.timestamp, table.callID, table.source, table.destination)
		assert.Nil(t, r, "Must be a invalid object!")
	}
}
