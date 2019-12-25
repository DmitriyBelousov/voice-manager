package finder

import "fmt"

//FinderIface ...
type FinderIface interface {
	OpenPhoneBook()
	FindContact() string
	ClosePhoneBook()
}

type finder struct{}

//OpenPhoneBook open phone contacts
func (c *finder) OpenPhoneBook() {
	fmt.Println("открываю телефонную книгу")
}

//FindContact find contact from voice command
func (c *finder) FindContact() string {
	fmt.Println("ищу контакт")
	return "Vasya"
}

//ClosePhoneBook close phone contacts
func (c *finder) ClosePhoneBook() {
	fmt.Println("закрываю телефонную книгу")
}

//NewFinder ...
func NewFinder() FinderIface {
	return &finder{}
}
