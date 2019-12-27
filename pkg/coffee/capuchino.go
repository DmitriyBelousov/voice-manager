package coffee

import (
	"github.com/DmitriyBelousov/voice-manager/pkg/models/coffeeMaker"
)

type Capuchino struct {
	coffee coffeeMaker.Coffee
}

func (b *Capuchino) SetWater() {
	b.coffee.Composition.WriteString("вода-")
}

func (b *Capuchino) SetCoffee() {
	b.coffee.Composition.WriteString("молотые зерна-")
}

func (b *Capuchino) SetAdditives() {
	b.coffee.Composition.WriteString("молоко")
}

func (b *Capuchino) GetProduct() coffeeMaker.Coffee {
	return b.coffee
}
