package findService

import "fmt"

type FinderIface interface {
	OpenPhoneBook()
	FindContact()
	ClosePhoneBook()
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

func (f *finderService) ClosePhoneBook() {
	fmt.Println("закрытие контактов")
}

func NewFinder(opt FinderOpts) FinderIface {
	return &finderService{
		Name: opt.Name,
	}
}
