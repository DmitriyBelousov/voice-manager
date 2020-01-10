package findService

import (
	"fmt"

	"github.com/pkg/errors"
)

type finder interface {
	OpenPhoneBook()
	FindContact(string) string
	ClosePoneBook()
}

// MobileFinder ...
type MobileFinder interface {
	OpenPhoneBook()
	FindContact(string) string
	ClosePhoneBook()
}

type finderService struct {
	name     string
	finder   finder
	contacts map[string]string
}

// OpenPhoneBook open phone contacts
func (f *finderService) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}

// FindContact find contact from voice command
func (f *finderService) FindContact(name string) string {
	fmt.Printf("поиск контакта, ищем в телефонной книге %s \n", name)

	return f.contacts[name]
}

// ClosePhoneBook close phone book
func (f *finderService) ClosePhoneBook() {
	fmt.Println("закрытие контактов")
}

// NewFinder ...
func NewFinder(cntcs map[string]string) (MobileFinder, error) {
	if cntcs != nil {
		return &finderService{contacts: cntcs}, nil
	}
	return nil, errors.New("create contact list from not initialized map")
}
