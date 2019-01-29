package interfaces

import (
	"fmt"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
	"github.com/jinzhu/gorm"
)

type DbStopCallRepo infrastructure.DbRepo

func NewDbStopCallRepo(conn *gorm.DB) (*DbStopCallRepo, error) {
	if conn := conn.AutoMigrate(&Record{}); conn.Error != nil {
		return nil, fmt.Errorf("failed to migrate record: %s", conn.Error.Error())
	}
	return &DbStopCallRepo{
		DB: conn,
	}, nil
}

func (repo *DbStopCallRepo) Store(recordStopCall domain.StopCall) error {
	if err := repo.DB.Create(&Record{
		Type:     2,
		RecordID: recordStopCall.RecordID,
		CallID:   recordStopCall.CallID,
	}); err.Error != nil {
		return fmt.Errorf("failed to create stop record: %s", err.Error.Error())
	}

	return nil
}

func (repo *DbStopCallRepo) FindById(id int) (*domain.StopCall, error) {
	var record Record

	if err := repo.DB.First(&record, id); err.Error != nil {
		return nil, fmt.Errorf("failed to find record: %s", err.Error.Error())
	}

	return &domain.StopCall{
		RecordID:  record.RecordID,
		Timestamp: record.CreatedAt,
		CallID:    record.CallID,
	}, nil
}
