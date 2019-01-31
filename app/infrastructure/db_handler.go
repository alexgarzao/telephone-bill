package infrastructure

import "github.com/jinzhu/gorm"

type DbHandler struct {
	*gorm.DB
}
