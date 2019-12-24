package facade

type manager struct {
	finder finder
	caller caller
	voicer voicer
}

type finder interface {
	OpenPhoneBook()
	FindContact()
	ClosePhoneBook()
}

type caller interface {
	MakeCall()
	CancelCall()
}

type voicer interface {
	ParseCommand()
	ParseName()
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
	m.finder.ClosePhoneBook()
}

type VoiceManager interface {
	Start()
	Stop()
}

func NewManager(f finder, c caller, v voicer) VoiceManager {
	return &manager{
		finder: f,
		caller: c,
		voicer: v,
	}
}
