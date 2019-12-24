package voiceService

import (
	"fmt"

	"github.com/DmitriyBelousov/voice-manager/pkg/models"
)

type voicer interface {
	ParseCommand()
	ParseName()
}

type VoicerIface interface {
	ParseCommand()
	ParseName()
}

type voicerService struct {
	Name   string
	voicer voicer
}

func (f *voicerService) ParseCommand() {
	fmt.Println("парсинг команды")
}

func (f *voicerService) ParseName() {
	fmt.Println("парсинг имени")
}

func NewVoicer(opt models.VoicerOpts) VoicerIface {
	return &voicerService{
		Name: opt.Name,
	}
}
