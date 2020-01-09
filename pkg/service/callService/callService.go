package callService

import (
	"fmt"

	models "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type caller interface {
	Call()
	CancelCall()
}

// MobileCaller ...
type MobileCaller interface {
	Call()
	CancelCall()
}

type callerService struct {
	name   string
	caller caller
}

// Call do call
func (f *callerService) Call() {
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
