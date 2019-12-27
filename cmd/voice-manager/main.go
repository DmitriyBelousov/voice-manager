package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	"github.com/DmitriyBelousov/voice-manager/pkg/models"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/callService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/findService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/voiceService"
	"github.com/DmitriyBelousov/voice-manager/pkg/user"
)

func main() {
	finder := findService.NewFinder(models.FinderOpts{Name: "name"})
	voicer := voiceService.NewVoicer(models.VoicerOpts{Name: "name"})
	caller := callService.NewCaller(models.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	userWithVoicer := user.NewUser(manager)
	userWithVoicer.UseVoiceManager()
	userWithVoicer.Sleep()
}
