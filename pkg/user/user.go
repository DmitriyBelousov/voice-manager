package user

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
)

type User interface {
	UseVoiceManager()
	Sleep()
}

type user struct {
	voiceManager facade.VoiceManager
}

func (u *user) UseVoiceManager() {
	fmt.Println("Достал телефон")

	u.voiceManager.Start()
	u.voiceManager.Stop()

	fmt.Println("Убрал телефон")
}

func (u *user) Sleep() {
	fmt.Println("Ушел спать")
}

func NewUser(manager facade.VoiceManager) User {
	return &user{
		voiceManager: manager,
	}
}
