package findService

import (
	"fmt"
	"testing"

	voice_manager "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

func Test_findService_NewFinder(t *testing.T) {
	name := "Vasya"
	caller := NewFinder(voice_manager.FinderOpts{Name: name})
	if res := caller.FindContact(); res != name {
		t.Error(fmt.Sprintf("Expected - %s", name))
	}

	if _, ok := caller.(FinderIface); !ok {
		t.Error("unexpected type")
	}
}
