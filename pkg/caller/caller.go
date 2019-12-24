package caller

import "fmt"

type Caller interface {
	MakeCall()
	CancelCall()
}

type CallerOpts struct {
	Name string
}

type caller struct {
	Name string
}

func (f *caller) MakeCall() {
	fmt.Println("делаем звонок")
}
func (f *caller) CancelCall() {
	fmt.Println("завершаем звонок")
}

func NewCaller(opt CallerOpts) Caller {
	return &caller{Name: opt.Name}
}
