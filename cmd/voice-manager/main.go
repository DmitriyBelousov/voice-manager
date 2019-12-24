package main

import (
	"github.com/DmitriyBelousov/voice-manager/internal/voiceManager"
	"github.com/DmitriyBelousov/voice-manager/user"
)

func main() {
	finder := voiceManager.NewFinder(voiceManager.FinderOpts{Name: "name"})
	voicer := voiceManager.NewVoicer(voiceManager.VoicerOpts{Name: "name"})
	caller := voiceManager.NewCaller(voiceManager.CallerOpts{Name: "name"})

	manager := voiceManager.NewManager(finder, caller, voicer)

	us := user.NewUser(manager)
	us.UseVoiceManager()
	us.Sleep()
}
