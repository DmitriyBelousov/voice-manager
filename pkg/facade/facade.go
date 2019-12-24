package facade

type finder interface {
	OpenPhoneBook()
	FindContact() string
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

type VoiceManager interface {
	Start()
	Stop()
}

type manager struct {
	finder finder
	caller caller
	voicer voicer
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

func NewManager(f finder, c caller, v voicer) VoiceManager {
	return &manager{
		finder: f,
		caller: c,
		voicer: v,
	}
}
