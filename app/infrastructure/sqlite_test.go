package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewSqliteWithValidArgument(t *testing.T) {
	db, err := NewSqlite(":memory:")
	assert.NotNil(t, db, "Must be a valid object!")
	assert.Nil(t, err, "Must be nil!")
}

func Test_NewSqliteWithInvalidArgument(t *testing.T) {
	db, err := NewSqlite("/a/b/c")
	assert.Nil(t, db, "Must be an invalid object!")
	assert.NotNil(t, err, "Must be not nil!")
	assert.Equal(t, err.Error(), "failed to connect database: unable to open database file")
}
