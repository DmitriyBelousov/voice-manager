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

// VoiceManager ...
type VoiceManager interface {
	Start()
	Stop()
}

type manager struct {
	finder finder
	caller caller
	voicer voicer
}

// Start make call action
func (m *manager) Start() {
	m.voicer.ParseCommand()
	m.voicer.ParseName()

	m.finder.OpenPhoneBook()
	m.finder.FindContact()

	m.caller.MakeCall()
}

// Finish call action
func (m *manager) Stop() {
	m.caller.CancelCall()
	m.finder.ClosePhoneBook()
}

// NewManager ...
func NewManager(finder finder, caller caller, voicer voicer) VoiceManager {
	return &manager{
		finder: finder,
		caller: caller,
		voicer: voicer,
	}
}
