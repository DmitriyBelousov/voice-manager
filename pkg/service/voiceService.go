package service

import "fmt"

type VoicerIface interface {
	ParseCommand()
	ParseName()
}

type voicer interface {
	ParseCommand()
	ParseName()
}

type VoicerOpts struct {
	Name string
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

func NewVoicer(opt VoicerOpts, v voicer) VoicerIface {
	return &voicerService{
		Name:   opt.Name,
		voicer: v,
	}
}
