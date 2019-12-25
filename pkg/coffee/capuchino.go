package coffee

import cm "github.com/DmitriyBelousov/voice-manager/pkg/coffeeMaker"

type Capuchino struct {
	coffee cm.Coffee
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

func (b *Capuchino) GetProduct() cm.Coffee {
	return b.coffee
}
