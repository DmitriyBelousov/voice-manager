package coffee

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

// Capuchino ...
type Capuchino struct {
	coffee coffeeMaker.Coffee
}

// SetWater ...
func (b *Capuchino) SetWater() {
	b.coffee.Composition.WriteString("вода-")
}

// SetCoffee ...
func (b *Capuchino) SetCoffee() {
	b.coffee.Composition.WriteString("молотые зерна-")
}

// SetAdditives ...
func (b *Capuchino) SetAdditives() {
	b.coffee.Composition.WriteString("молоко")
}

// SetAdditives ...
func (b *Capuchino) GetProduct() coffeeMaker.Coffee {
	return b.coffee
}
