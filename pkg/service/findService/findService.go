package findService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type finder interface {
	OpenPhoneBook()
	FindContact() string
	ClosePoneBook()
}

// FinderIface ...
type FinderIface interface {
	OpenPhoneBook()
	FindContact() string
	ClosePhoneBook()
}

type finderService struct {
	Name   string
	finder finder
}

// OpenPhoneBook open phone contacts
func (f *finderService) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}

// FindContact find contact from voice command
func (f *finderService) FindContact() string {
	fmt.Println("поиск контакта")
	return "Vasya"
}

// ClosePhoneBook close phone book
func (f *finderService) ClosePhoneBook() {
	fmt.Println("закрытие контактов")
}

// NewFinder ...
func NewFinder(opt models.FinderOpts) FinderIface {
	return &finderService{
		Name: opt.Name,
	}
}
