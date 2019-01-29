package usecases

import (
	"testing"
	"time"

	"github.com/alexgarzao/telephone-bill/app/infrastructure"
	"github.com/alexgarzao/telephone-bill/app/interfaces"
	"github.com/stretchr/testify/assert"
)

func TestValidRecordStartCall(t *testing.T) {
	dbHandler, _ := infrastructure.NewSqlite(":memory:")

	r := new(RecordStartCallInteractor)
	r.StartCallRepository, _ = interfaces.NewDbStartCallRepo(dbHandler)
	assert.Nil(t, r.Add("R1", time.Now(), "C1", "1212345678", "12123456789"), "Must be a valid object!")
}

func TestInvalidRecordStartCall(t *testing.T) {
	dbHandler, _ := infrastructure.NewSqlite(":memory:")

	r := new(RecordStartCallInteractor)
	r.StartCallRepository, _ = interfaces.NewDbStartCallRepo(dbHandler)
	r.Logger = new(infrastructure.Logger)
	assert.NotNil(t, r.Add("R1", time.Now(), "C1", "12123", "12123456789"), "Must be an invalid object!")
}
