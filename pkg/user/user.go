package user

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
)

//User ...
type user struct {
	voiceManager facade.VoiceManager
}

type User interface {
	UseVoiceManager()
	Sleep()
}

//New ...
func NewUser(manager facade.VoiceManager) User {
	return &user{
		voiceManager: manager,
	}
}

//UseVoiceManager ...
func (u *user) UseVoiceManager() {
	fmt.Println("Достал телефон")

	u.voiceManager.Start()
	u.voiceManager.Stop()

	fmt.Println("Убрал телефон")
}

func (u *user) Sleep() {
	fmt.Println("Ушел спать")
}
