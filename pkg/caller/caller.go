package caller

import "fmt"

type CallerIface interface {
	MakeCall()
	CancelCall()
}

type caller struct{}

func (c *caller) MakeCall() {
	fmt.Println("делаю звонок")
}

func (c *caller) CancelCall() {
	fmt.Println("отменяю звонок")
}

func NewCaller() CallerIface {
	return &caller{}
}
