package facade

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/caller"
	"github.com/DmitriyBelousov/voice-manager/pkg/finder"
	"github.com/DmitriyBelousov/voice-manager/pkg/voicer"
)

type manager struct {
	finder finder.Finder
	caller caller.Caller
	voicer voicer.Voicer
}

func (m *manager) Start() {
	m.voicer.ParseCommand()
	m.voicer.ParseName()

	m.finder.OpenPhoneBook()
	m.finder.FindContact()

	m.caller.MakeCall()
}
func (m *manager) Stop() {
	m.caller.CancelCall()
	m.finder.ClosePoneBook()
}

type VoiceManager interface {
	Start()
	Stop()
}

//New factory func
func NewManager(f finder.Finder, c caller.Caller, v voicer.Voicer) VoiceManager {
	return &manager{
		finder: f,
		caller: c,
		voicer: v,
	}
}
