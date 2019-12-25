package coffee

import (
	cm "github.com/DmitriyBelousov/voice-manager/pkg/coffeeMaker"
)

type Americano struct {
	coffee cm.Coffee
}

func (b *Americano) SetWater() {
	b.coffee.Composition.WriteString("вода-")
}

func (b *Americano) SetCoffee() {
	b.coffee.Composition.WriteString("молотые зерна-")
}

func (b *Americano) SetAdditives() {
}

func (b *Americano) GetProduct() cm.Coffee {
	return b.coffee
}
