package service

import "fmt"

type FinderIface interface {
	OpenPhoneBook()
	FindContact()
	ClosePoneBook()
}

type finder interface {
	OpenPhoneBook()
	FindContact()
	ClosePoneBook()
}

type FinderOpts struct {
	Name string
}

type finderService struct {
	Name   string
	finder finder
}

func (f *finderService) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}
func (f *finderService) FindContact() {
	fmt.Println("поиск контакта")
}

func (f *finderService) ClosePoneBook() {
	fmt.Println("закрытие контактов")
}

func NewFinder(opt FinderOpts, f finder) FinderIface {
	return &finderService{
		Name:   opt.Name,
		finder: f,
	}
}
