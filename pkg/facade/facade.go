package voiceManager

type manager struct {
	finder Finder
	caller Caller
	voicer Voicer
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
func NewManager(f Finder, c Caller, v Voicer) VoiceManager {
	return &manager{
		finder: f,
		caller: c,
		voicer: v,
	}
}
