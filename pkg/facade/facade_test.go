package facade

import (
	"testing"

	"git.wildberries.ru/portals/feedback-service/_vendor-20191224161050/github.com/stretchr/testify/assert"
)

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
	_, ok := manager.(VoiceManager)
	assert.Equal(t, true, ok)

}
