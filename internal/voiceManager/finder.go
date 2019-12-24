package voiceManager

type Finder interface{
	OpenPhoneBook()
	FindContact()
	ClosePoneBook()
}