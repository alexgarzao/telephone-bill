package interfaces

import (
	"fmt"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
	"github.com/jinzhu/gorm"
)

type DbStartCallRepo infrastructure.DbHandler

func NewDbStartCallRepo(conn *gorm.DB) (*DbStartCallRepo, error) {
	if conn := conn.AutoMigrate(&Record{}); conn.Error != nil {
		return nil, fmt.Errorf("failed to migrate record: %s", conn.Error.Error())
	}
	return &DbStartCallRepo{
		DB: conn,
	}, nil
}

func (repo *DbStartCallRepo) Store(recordStartCall domain.StartCall) error {
	if err := repo.DB.Create(&Record{
		Type:        1,
		RecordID:    recordStartCall.RecordID,
		CallID:      recordStartCall.CallID,
		Source:      recordStartCall.Source,
		Destination: recordStartCall.Destination,
	}); err.Error != nil {
		return fmt.Errorf("failed to create start record: %s", err.Error.Error())
	}

	return nil
}

func (repo *DbStartCallRepo) FindById(id int) (*domain.StartCall, error) {
	var record Record

	if err := repo.DB.First(&record, id); err.Error != nil {
		return nil, fmt.Errorf("failed to find record: %s", err.Error.Error())
	}

	return &domain.StartCall{
		RecordID:    record.RecordID,
		Timestamp:   record.CreatedAt,
		CallID:      record.CallID,
		Source:      record.Source,
		Destination: record.Destination,
	}, nil
}
