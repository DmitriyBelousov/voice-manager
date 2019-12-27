package findService

import (
	"fmt"

	voice_manager "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type finder interface {
	OpenPhoneBook()
	FindContact() string
	ClosePoneBook()
}

//FinderIface ...
type FinderIface interface {
	OpenPhoneBook()
	FindContact() string
	ClosePhoneBook()
}

type finderService struct {
	Name   string
	finder finder
}

//OpenPhoneBook open phone contacts
func (f *finderService) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}

//FindContact find contact from voice command
func (f *finderService) FindContact() string {
	fmt.Println("поиск контакта")
	return "Vasya"
}

//ClosePhoneBook close phone book
func (f *finderService) ClosePhoneBook() {
	fmt.Println("закрытие контактов")
}

//NewFinder ...
func NewFinder(opt voice_manager.FinderOpts) FinderIface {
	return &finderService{
		Name: opt.Name,
	}
}
