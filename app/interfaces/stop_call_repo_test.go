package interfaces

import (
	_ "fmt"
	"testing"
	"time"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
	"github.com/stretchr/testify/assert"
)

func Test_ValidStopCallRepository(t *testing.T) {
	dbHandler, _ := infrastructure.NewSqlite(":memory:")
	scr, _ := NewDbStopCallRepo(dbHandler)

	now := time.Now()

	sc11 := domain.StopCall{
		RecordID:  "r1",
		Timestamp: now,
		CallID:    "c1",
	}

	scr.Store(sc11)

	sc12, _ := scr.FindById(1)
	assert.Equal(t, sc11.RecordID, sc12.RecordID)
	assert.Equal(t, sc11.Timestamp, now)
	assert.Equal(t, sc11.CallID, sc12.CallID)
}

func Test_InvalidIdInStopCallRepository(t *testing.T) {
	dbHandler, _ := infrastructure.NewSqlite(":memory:")
	scr, _ := NewDbStopCallRepo(dbHandler)

	now := time.Now()

	sc11 := domain.StopCall{
		RecordID:  "r1",
		Timestamp: now,
		CallID:    "c1",
	}

	scr.Store(sc11)
	sc12, err := scr.FindById(10)
	assert.Nil(t, sc12, "Must be nil")
	assert.Equal(t, err.Error(), "failed to find record: record not found")
}
