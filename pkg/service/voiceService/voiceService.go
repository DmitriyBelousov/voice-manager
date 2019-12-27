package voiceService

import (
	"fmt"

	voice_manager "github.com/DmitriyBelousov/voice-manager/pkg/models/voice-manager"
)

type voicer interface {
	ParseCommand()
	ParseName()
}

//VoicerIface ...
type VoicerIface interface {
	ParseCommand()
	ParseName()
}

type voicerService struct {
	Name   string
	voicer voicer
}

//ParseCommand get action from voice command
func (f *voicerService) ParseCommand() {
	fmt.Println("парсинг команды")
}

//ParseName get name param for command
func (f *voicerService) ParseName() {
	fmt.Println("парсинг имени")
}

//NewVoicer ...
func NewVoicer(opt voice_manager.VoicerOpts) VoicerIface {
	return &voicerService{
		Name: opt.Name,
	}
}
