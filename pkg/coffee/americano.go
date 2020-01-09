package coffee

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

// Americano ...
type Americano struct {
	coffee coffeeMaker.Coffee
}

// SetWater ...
func (b *Americano) SetWater() {
	b.coffee.Composition.WriteString("вода-")
}

// SetCoffee ...
func (b *Americano) SetCoffee() {
	b.coffee.Composition.WriteString("молотые зерна-")
}

// SetAdditives ...
func (b *Americano) SetAdditives() {
}

// GetProduct ...
func (b *Americano) GetProduct() coffeeMaker.Coffee {
	return b.coffee
}
