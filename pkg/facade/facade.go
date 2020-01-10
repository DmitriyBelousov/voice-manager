package facade

type finder interface {
	OpenPhoneBook()
	FindContact(string) string
	ClosePhoneBook()
}

type caller interface {
	Call(string)
	CancelCall()
}

type voicer interface {
	ParseCommand()
	ParseName(string) string
}

// VoiceManager ...
type VoiceManager interface {
	Start(string)
	Stop()
}

type manager struct {
	finder finder
	caller caller
	voicer voicer
}

// Start make call action
func (m *manager) Start(name string) {
	m.voicer.ParseCommand()
	parsedName := m.voicer.ParseName(name)

	m.finder.OpenPhoneBook()
	number := m.finder.FindContact(parsedName)

	m.caller.Call(number)
}

// Stop finish call action
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
