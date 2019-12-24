package facade

import (
	"testing"

	"github.com/DmitriyBelousov/voice-manager/pkg/service/callService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/findService"
	"github.com/DmitriyBelousov/voice-manager/pkg/service/voiceService"
)

func Test_facade_NewManager(t *testing.T) {
	finder := findService.NewFinder(findService.FinderOpts{Name: "name"})
	voicer := voiceService.NewVoicer(voiceService.VoicerOpts{Name: "name"})
	caller := callService.NewCaller(callService.CallerOpts{Name: "name"})

	manager := NewManager(finder, caller, voicer)
	if _, ok := manager.(VoiceManager); !ok {
		t.Error("unexpected type")
	}
}
