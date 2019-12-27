package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	voice_manager "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/callService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/findService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/voiceService"
)

func main() {
	finder := findService.NewFinder(voice_manager.FinderOpts{Name: "name"})
	voicer := voiceService.NewVoicer(voice_manager.VoicerOpts{Name: "name"})
	caller := callService.NewCaller(voice_manager.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	userWithVoicer := facade.NewUser(manager)
	userWithVoicer.UseVoiceManager()
	userWithVoicer.Sleep()
}
