package callService

import (
	"fmt"

	models "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type caller interface {
	MakeCall()
	CancelCall()
}

// MobileCaller ...
type MobileCaller interface {
	MakeCall()
	CancelCall()
}

type callerService struct {
	name   string
	caller caller
}

// MakeCall do call
func (f *callerService) MakeCall() {
	fmt.Println("делаем звонок")
}

// CancelCall cancel call
func (f *callerService) CancelCall() {
	fmt.Println("завершаем звонок")
}

// NewCaller ...
func NewCaller(opt models.CallerOpts) MobileCaller {
	return &callerService{
		name: opt.Name,
	}
}
