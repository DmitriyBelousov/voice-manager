package voiceManager

import "fmt"

type Voicer interface {
	ParseCommand()
	ParseName()
}

type VoicerOpts struct {
	Name string
}

type voicer struct {
	Name string
}

func (f *voicer) ParseCommand() {
	fmt.Println("парсинг команды")
}
func (f *voicer) ParseName() {
	fmt.Println("парсинг имени")
}

func NewVoicer(opt VoicerOpts) Voicer {
	return &voicer{Name: opt.Name}
}
