package findService

import (
	"fmt"
	"testing"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

func Test_findService_NewFinder(t *testing.T) {
	name := "Vasya"
	caller := NewFinder(models.FinderOpts{Name: name})
	if res := caller.FindContact(); res != name {
		t.Error(fmt.Sprintf("Expected - %s", name))
	}

	if _, ok := caller.(MobileFinder); !ok {
		t.Error("unexpected type")
	}
}
