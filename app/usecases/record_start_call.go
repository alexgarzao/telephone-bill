package usecases

import (
	"fmt"
	"time"

	"github.com/alexgarzao/telephone-bill/app/domain"
	"github.com/alexgarzao/telephone-bill/app/infrastructure"
)

type RecordStartCallInteractor struct {
	StartCallRepository domain.StartCallRepository
	Logger              *infrastructure.Logger
}

func (interactor *RecordStartCallInteractor) Add(recordID string, timestamp time.Time, callID string, source string, destination string) error {
	r, err := domain.NewStartCall(recordID, timestamp, callID, source, destination)
	if err != nil {
		err := fmt.Errorf("Impossible to add Start Call: %s", err.Error())
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.StartCallRepository.Store(*r)

	return nil
}
