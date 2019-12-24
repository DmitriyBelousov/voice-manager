package finder

import "fmt"

type FinderIface interface {
	OpenPhoneBook()
	FindContact()
	ClosePhoneBook()
}

type finder struct{}

func (c *finder) OpenPhoneBook() {
	fmt.Println("открываю телефонную книгу")
}

func (c *finder) FindContact() {
	fmt.Println("ищу контакт")
}

func (c *finder) ClosePhoneBook() {
	fmt.Println("закрываю телефонную книгу")
}

func NewFinder() FinderIface {
	return &finder{}
}
