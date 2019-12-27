package coffee

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

type Americano struct {
	coffee coffeeMaker.Coffee
}

func (b *Americano) SetWater() {
	b.coffee.Composition.WriteString("вода-")
}

func (b *Americano) SetCoffee() {
	b.coffee.Composition.WriteString("молотые зерна-")
}

func (b *Americano) SetAdditives() {
}

func (b *Americano) GetProduct() coffeeMaker.Coffee {
	return b.coffee
}
