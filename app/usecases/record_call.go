package usecases

import (
	"fmt"
	"time"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
)

type RecordCallInteractor struct {
	StartCallRepository domain.StartCallRepository
	StopCallRepository  domain.StopCallRepository
	Logger              *infrastructure.Logger
}

func (interactor *RecordCallInteractor) AddStart(recordID string, timestamp time.Time, callID string, source string, destination string) error {
	r, err := domain.NewStartCall(recordID, timestamp, callID, source, destination)
	if err != nil {
		err := fmt.Errorf("Impossible to add Start Call: %s", err.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.StartCallRepository.Store(*r)

	return nil
}

func (interactor *RecordCallInteractor) AddStop(recordID string, timestamp time.Time, callID string) error {
	r, err := domain.NewStopCall(recordID, timestamp, callID)
	if err != nil {
		err := fmt.Errorf("Impossible to add Stop Call: %s", err.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.StopCallRepository.Store(*r)

	return nil
}
