package voicer

import "fmt"

// MobileVoicer ...
type MobileVoicer interface {
	ParseCommand()
	ParseName()
}

type voicer struct{}

// ParseCommand get action from voice command
func (c *voicer) ParseCommand() {
	fmt.Println("распознаю команду")
}

// ParseName get name param for command
func (c *voicer) ParseName() {
	fmt.Println("распознаю имя")
}

// NewVoicer ...
func NewVoicer() MobileVoicer {
	return &voicer{}
}
