package facade

import "testing"

type mockFinder struct {
}
type mockCaller struct {
}
type mockVoicer struct {
}

func (mf *mockFinder) OpenPhoneBook()      {}
func (mf *mockFinder) FindContact() string { return "" }
func (mf *mockFinder) ClosePhoneBook()     {}

func (mc *mockCaller) MakeCall()   {}
func (mc *mockCaller) CancelCall() {}

func (mv *mockVoicer) ParseCommand() {}
func (mv *mockVoicer) ParseName()    {}

func Test_facade_NewManager(t *testing.T) {

	f := new(mockFinder)
	c := new(mockCaller)
	v := new(mockVoicer)

	manager := NewManager(f, c, v)
	if _, ok := manager.(VoiceManager); !ok {
		t.Error("unexpected type")
	}
}
