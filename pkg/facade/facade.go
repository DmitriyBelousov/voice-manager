package facade

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/service"
)

type manager struct {
	finder service.Finder
	caller service.Caller
	voicer service.Voicer
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

func NewManager(f service.Finder, c service.Caller, v service.Voicer) VoiceManager {
	return &manager{
		finder: f,
		caller: c,
		voicer: v,
	}
}
