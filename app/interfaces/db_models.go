package interfaces

import (
	"github.com/jinzhu/gorm"
)

type Record struct {
	// TODO: is possible to use start_call from domain????
	gorm.Model
	RecordID    string
	Type        uint
	CallID      string
	Source      string
	Destination string
}
