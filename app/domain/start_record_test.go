package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidStartRecord(t *testing.T) {
	r := NewStartRecord("R1", time.Now(), "C1", "S1", "D1")
	assert.NotNil(t, r, "Must be a valid object!")
}
