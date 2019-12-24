package main

import (
	c "github.com/DmitriyBelousov/voice-manager/pkg/caller"
	"github.com/DmitriyBelousov/voice-manager/pkg/facade"
	f "github.com/DmitriyBelousov/voice-manager/pkg/finder"
	"github.com/DmitriyBelousov/voice-manager/pkg/user"
	v "github.com/DmitriyBelousov/voice-manager/pkg/voicer"
)

func main() {
	finder := f.NewFinder(f.FinderOpts{Name: "name"})
	voicer := v.NewVoicer(v.VoicerOpts{Name: "name"})
	caller := c.NewCaller(c.CallerOpts{Name: "name"})

	manager := facade.NewManager(finder, caller, voicer)

	us := user.NewUser(manager)
	us.UseVoiceManager()
	us.Sleep()
}
