package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStartCall(t *testing.T) {
	n := time.Now()

	tables := []StartCall{
		{"R1", n, "C1", "1212345678", "12123456789"},
		{"R1123", n, "asasdasd", "1232345678", "12123456789"},
		{"999", n, "23123132@#@!@@", "1212222678", "12123456789"},
		{"aaAAA321312", n, "C1", "12123456222", "12123456789"},
	}

	for _, table := range tables {
		r, _ := NewStartCall(table.RecordID, table.Timestamp, table.CallID, table.Source, table.Destination)
		assert.NotNil(t, r, "Must be a valid object!")
		assert.Equal(t, r.RecordID, table.RecordID)
		assert.Equal(t, r.Timestamp, table.Timestamp)
		assert.Equal(t, r.CallID, table.CallID)
		assert.Equal(t, r.Source, table.Source)
		assert.Equal(t, r.Destination, table.Destination)
	}
}

func TestInvalidStartCallWhenSomeFieldIsEmpty(t *testing.T) {
	n := time.Now()

	tables := []StartCall{
		{"", n, "C1", "1212345678", "12123456789"},
		{"R1", time.Time{}, "C1", "1212345678", "12123456789"},
		{"R1", n, "", "1212345678", "12123456789"},
		{"R1", n, "C1", "", "12123456789"},
		{"R1", n, "C1", "1212345678", ""},
	}

	for _, table := range tables {
		r, err := NewStartCall(table.RecordID, table.Timestamp, table.CallID, table.Source, table.Destination)
		assert.Nil(t, r, "Must be a invalid object!")
		assert.EqualError(t, err, "empty fields")
	}
}

func TestInvalidStartCallWhenSomeFieldIsInvalid(t *testing.T) {
	n := time.Now()

	tables := []StartCall{
		{"R1", n, "C1", "12145678", "12123456789"},
		{"R1", n, "C1", "1212345678", "1212789"},
		{"R1", n, "C1", "1212345678", "1212789123123"},
	}

	errors := []string{
		"invalid source number: 12145678",
		"invalid destination number: 1212789",
		"invalid destination number: 1212789123123",
	}
	errorIdx := 0

	for _, table := range tables {
		r, err := NewStartCall(table.RecordID, table.Timestamp, table.CallID, table.Source, table.Destination)
		assert.Nil(t, r, "Must be a invalid object!")
		assert.EqualError(t, err, errors[errorIdx])
		errorIdx++
	}
}
