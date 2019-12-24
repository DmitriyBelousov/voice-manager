package voicer

import "fmt"

type VoicerIface interface {
	ParseCommand()
	ParseName()
}

type voicer struct{}

func (c *voicer) ParseCommand() {
	fmt.Println("распознаю команду")
}

func (c *voicer) ParseName() {
	fmt.Println("распознаю имя")
}

func NewVoicer() VoicerIface {
	return &voicer{}
}
