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

type FinderIface interface {
	OpenPhoneBook()
	FindContact() string
	ClosePhoneBook()
}

type finderService struct {
	Name   string
	finder finder
}

func (f *finderService) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}

func (f *finderService) FindContact() string {
	fmt.Println("поиск контакта")
	return "Vasya"
}

func (f *finderService) ClosePhoneBook() {
	fmt.Println("закрытие контактов")
}

func NewFinder(opt models.FinderOpts) FinderIface {
	return &finderService{
		Name: opt.Name,
	}
}
