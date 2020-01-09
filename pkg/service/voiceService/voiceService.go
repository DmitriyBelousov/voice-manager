package voiceService

import (
	"fmt"

	models "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type voicer interface {
	ParseCommand()
	ParseName()
}

// Voicer ...
type Voicer interface {
	ParseCommand()
	ParseName()
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
func (f *voicerService) ParseName() {
	fmt.Println("парсинг имени")
}

// NewVoicer ...
func NewVoicer(opt models.VoicerOpts) Voicer {
	return &voicerService{
		name: opt.Name,
	}
}
