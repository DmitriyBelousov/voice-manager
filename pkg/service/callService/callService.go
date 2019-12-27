package callService

import (
	"fmt"

	voice_manager "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type caller interface {
	MakeCall()
	CancelCall()
}

//CallerIface ...
type CallerIface interface {
	MakeCall()
	CancelCall()
}

type callerService struct {
	Name   string
	caller caller
}

//MakeCall do call
func (f *callerService) MakeCall() {
	fmt.Println("делаем звонок")
}

//CancelCall cancel call
func (f *callerService) CancelCall() {
	fmt.Println("завершаем звонок")
}

//NewCaller ...
func NewCaller(opt voice_manager.CallerOpts) CallerIface {
	return &callerService{
		Name: opt.Name,
	}
}
