package callService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type caller interface {
	Call(string)
	CancelCall()
}

// MobileCaller ...
type MobileCaller interface {
	Call(string)
	CancelCall()
}

type callerService struct {
	name   string
	caller caller
}

// Call do call
func (f *callerService) Call(number string) {
	fmt.Printf("делаем звонок на номер %s \n", number)
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
