package user

import (
	"fmt"
)

type voiceManager interface {
	Start(string)
	Stop()
}

// User ...
type User interface {
	UseVoiceManager(string)
	Sleep()
}

type user struct {
	voiceManager voiceManager
}

// UseVoiceManager facade usage
func (u *user) UseVoiceManager(name string) {
	fmt.Println("Достал телефон")

	u.voiceManager.Start(name)
	u.voiceManager.Stop()

	fmt.Println("Убрал телефон")
}

// Sleep ...
func (u *user) Sleep() {
	fmt.Println("Ушел спать")
}

// NewUser ...
func NewUser(manager voiceManager) User {
	return &user{
		voiceManager: manager,
	}
}
