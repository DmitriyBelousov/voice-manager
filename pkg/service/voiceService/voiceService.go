package voiceService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
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
func NewVoicer(opt models.VoicerOpts) VoicerIface {
	return &voicerService{
		Name: opt.Name,
	}
}
