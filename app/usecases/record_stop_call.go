package usecases

import (
	"fmt"
	"time"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
)

type RecordStopCallInteractor struct {
	StopCallRepository domain.StopCallRepository
	Logger             *infrastructure.Logger
}

func (interactor *RecordStopCallInteractor) Add(recordID string, timestamp time.Time, callID string) error {
	r, err := domain.NewStopCall(recordID, timestamp, callID)
	if err != nil {
		err := fmt.Errorf("Impossible to add Stop Call: %s", err.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.StopCallRepository.Store(*r)

	return nil
}
