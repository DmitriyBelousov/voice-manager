package finder

import "fmt"

type Finder interface {
	OpenPhoneBook()
	FindContact()
	ClosePoneBook()
}

type FinderOpts struct {
	Name string
}

type finder struct {
	Name string
}

func (f *finder) OpenPhoneBook() {
	fmt.Println("открытие контактов")
}
func (f *finder) FindContact() {
	fmt.Println("поиск контакта")
}

func (f *finder) ClosePoneBook() {
	fmt.Println("закрытие контактов")
}

func NewFinder(opt FinderOpts) Finder {
	return &finder{Name: opt.Name}
}
