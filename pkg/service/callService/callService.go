package callService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type caller interface {
	MakeCall()
	CancelCall()
}

type CallerIface interface {
	MakeCall()
	CancelCall()
}

type callerService struct {
	Name   string
	caller caller
}

func (f *callerService) MakeCall() {
	fmt.Println("делаем звонок")
}
func (f *callerService) CancelCall() {
	fmt.Println("завершаем звонок")
}

func NewCaller(opt models.CallerOpts) CallerIface {
	return &callerService{
		Name: opt.Name,
	}
}
