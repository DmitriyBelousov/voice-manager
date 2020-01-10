package voiceService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type voicer interface {
	ParseCommand()
	ParseName(string) string
}

// Voicer ...
type Voicer interface {
	ParseCommand()
	ParseName(string) string
}

type voicerService struct {
	name   string
	voicer voicer
}

// ParseCommand get action from voice command
func (f *voicerService) ParseCommand() {
	fmt.Println("парсинг команды")
}

// ParseName get name param for command
func (f *voicerService) ParseName(name string) string {
	fmt.Printf("парсинг имени, ищу %s \n", name)

	return name
}

// NewVoicer ...
func NewVoicer(opt models.VoicerOpts) Voicer {
	return &voicerService{
		name: opt.Name,
	}
}
