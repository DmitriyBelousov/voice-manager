package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	"github.com/DmitriyBelousov/voice-manager/pkg/models"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/callService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/findService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/voiceService"
	"github.com/DmitriyBelousov/voice-manager/pkg/user"
	"github.com/rs/zerolog/log"
)

func main() {
	contacts := map[string]string{
		"Vasya": "8-904-235-44-33",
		"Petay": "8-907-235-55-34",
	}

	finder, err := findService.NewFinder(contacts)
	if err != nil {
		log.Error().Msg("finder init error")
		return
	}

	voicer := voiceService.NewVoicer(models.VoicerOpts{Name: "name"})
	caller := callService.NewCaller(models.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	userWithVoicer := user.NewUser(manager)
	userWithVoicer.UseVoiceManager("Petay")
	userWithVoicer.UseVoiceManager("Vasya")

	userWithVoicer.Sleep()
}
