package user

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
)

//User ...
type User interface {
	UseVoiceManager()
	Sleep()
}

type user struct {
	voiceManager facade.VoiceManager
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
func NewUser(manager facade.VoiceManager) User {
	return &user{
		voiceManager: manager,
	}
}
