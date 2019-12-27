package callService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type caller interface {
	MakeCall()
	CancelCall()
}

// CallerIface ...
type CallerIface interface {
	MakeCall()
	CancelCall()
}

type callerService struct {
	Name   string
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
func NewCaller(opt models.CallerOpts) CallerIface {
	return &callerService{
		Name: opt.Name,
	}
}
