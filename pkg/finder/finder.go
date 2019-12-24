package finder

import "fmt"

type FinderIface interface {
	OpenPhoneBook()
	FindContact() string
	ClosePhoneBook()
}

type finder struct{}

func (c *finder) OpenPhoneBook() {
	fmt.Println("открываю телефонную книгу")
}

func (c *finder) FindContact() string {
	fmt.Println("ищу контакт")
	return "Vasya"
}

func (c *finder) ClosePhoneBook() {
	fmt.Println("закрываю телефонную книгу")
}

func NewFinder() FinderIface {
	return &finder{}
}
