package service

import "fmt"

type caller interface {
	MakeCall()
	CancelCall()
}

type CallerIface interface {
	MakeCall()
	CancelCall()
}

type CallerOpts struct {
	Name string
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

func NewCaller(opt CallerOpts, c caller) CallerIface {
	return &callerService{
		Name:   opt.Name,
		caller: c,
	}
}
