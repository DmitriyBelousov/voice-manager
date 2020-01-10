package finder

import (
	"errors"
	"fmt"
)

// MobileFinder ...
type MobileFinder interface {
	OpenPhoneBook()
	FindContact(string) string
	ClosePhoneBook()
}

type finder struct {
	contacts map[string]string
}

// OpenPhoneBook open phone contacts
func (f *finder) OpenPhoneBook() {
	fmt.Println("открываю телефонную книгу")
}

// FindContact find contact from voice command
func (f *finder) FindContact(name string) string {
	fmt.Printf("парсинг имени, ищу %s \n", name)

	return f.contacts[name]
}

// ClosePhoneBook close phone contacts
func (f *finder) ClosePhoneBook() {
	fmt.Println("закрываю телефонную книгу")
}

// NewFinder ...
func NewFinder(cntcs map[string]string) (MobileFinder, error) {
	if cntcs != nil {
		return &finder{contacts: cntcs}, nil

	}
	return nil, errors.New("create contact list from not initialized map")
}
