package findService

import (
	"testing"

	"git.wildberries.ru/portals/feedback-service/_vendor-20191224161050/github.com/stretchr/testify/assert"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

func Test_findService_NewFinder(t *testing.T) {
	actual := "Vasya"
	caller := NewFinder(models.FinderOpts{Name: actual})
	res := caller.FindContact()

	assert.Equal(t, actual, res)

}
