package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/callService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/findService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/voiceService"
	"github.com/DmitriyBelousov/voice-manager/pkg/user"
)

func main() {
	finder := findService.NewFinder(findService.FinderOpts{Name: "name"})
	voicer := voiceService.NewVoicer(voiceService.VoicerOpts{Name: "name"})
	caller := callService.NewCaller(callService.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	userWithVoicer := user.NewUser(manager)
	userWithVoicer.UseVoiceManager()
	userWithVoicer.Sleep()
}
