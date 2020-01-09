package visitor

type passengerVisitor interface {
	visitBusinessPassenger(bp *businessPassenger)
	visitEconomPassenger(ep *economPassenger)
}

type Passenger interface {
	Accept(passengerVisitor)
}

type businessPassenger struct {
	Name string
}

func (bp *businessPassenger) Accept(visitor passengerVisitor) {
	visitor.visitBusinessPassenger(bp)
}

type economPassenger struct {
	Name string
}

func (this *economPassenger) Accept(visitor passengerVisitor) {
	visitor.visitEconomPassenger(this)
}

type PassengersList struct {
	passengers []Passenger
}

func (pl *PassengersList) Add(pass Passenger) {
	pl.passengers = append(pl.passengers, pass)
}

func (pl *PassengersList) Accept(visitor passengerVisitor) {
	for _, pass := range pl.passengers {
		pass.Accept(visitor)
	}
}

func NewBusinessPassenger(name string) *businessPassenger {
	return &businessPassenger{Name: name}
}

func NewEconomPassenger(name string) *economPassenger {
	return &economPassenger{Name: name}
}
