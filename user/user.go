package user

import (
	"fmt"

	vm "github.com/DmitriyBelousov/voice-manager/internal/voiceManager"
)

var facade vm.VoiceManager

//User ...
type user struct {
	voiceManager vm.VoiceManager
}

type User interface {
	UseVoiceManager()
	Sleep()
}

//New ...
func NewUser(manager vm.VoiceManager) User {
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
