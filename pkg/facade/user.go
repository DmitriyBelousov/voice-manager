package facade

import (
	"fmt"
)

//User ...
type User interface {
	UseVoiceManager()
	Sleep()
}

type user struct {
	voiceManager VoiceManager
}

//UseVoiceManager facade usage
func (u *user) UseVoiceManager() {
	fmt.Println("Достал телефон")

	u.voiceManager.Start()
	u.voiceManager.Stop()

	fmt.Println("Убрал телефон")
}

//Sleep ...
func (u *user) Sleep() {
	fmt.Println("Ушел спать")
}

//NewUser ...
func NewUser(manager VoiceManager) User {
	return &user{
		voiceManager: manager,
	}
}
