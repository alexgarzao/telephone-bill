package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func NewSqlite(dbfileName string) (*gorm.DB, error) {
	conn, err := gorm.Open("sqlite3", dbfileName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %s", err.Error())
	}
	return conn, nil
}
