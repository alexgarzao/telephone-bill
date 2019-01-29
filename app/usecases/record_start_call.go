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
	r := domain.NewStartCall(recordID, timestamp, callID, source, destination)
	if r == nil {
		err := fmt.Errorf("Impossible to add Start Call: recordID %s timestamp %s callID %s source %s destination %s",
			recordID, timestamp, callID, source, destination,
		)
		interactor.Logger.Log(err.Error())
		return err
	}
	interactor.StartCallRepository.Store(*r)

	return nil
}
