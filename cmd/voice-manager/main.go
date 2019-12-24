package main

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	"github.com/DmitriyBelousov/voice-manager/pkg/service"
	"github.com/DmitriyBelousov/voice-manager/pkg/user"
)

func main() {
	finder := service.NewFinder(service.FinderOpts{Name: "name"})
	voicer := service.NewVoicer(service.VoicerOpts{Name: "name"})
	caller := service.NewCaller(service.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	us := user.NewUser(manager)
	us.UseVoiceManager()
	us.Sleep()
}
