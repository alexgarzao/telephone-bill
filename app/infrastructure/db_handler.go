package infrastructure

import "github.com/jinzhu/gorm"

type DbRepo struct {
	*gorm.DB
}
