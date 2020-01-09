package caller

import "fmt"

// MobileCaller ...
type MobileCaller interface {
	Call()
	CancelCall()
}

type caller struct{}

// Call do call
func (c *caller) Call() {
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
