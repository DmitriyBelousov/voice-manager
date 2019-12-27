package caller

import "fmt"

// MobileCaller ...
type MobileCaller interface {
	MakeCall()
	CancelCall()
}

type caller struct{}

// MakeCall do call
func (c *caller) MakeCall() {
	fmt.Println("делаю звонок")
}

// CancelCall cancel call
func (c *caller) CancelCall() {
	fmt.Println("отменяю звонок")
}

// NewCaller ...
func NewCaller() MobileCaller {
	return &caller{}
}
