package visitor

type passengerVisitor interface {
	visitBusinessPassenger(bp *businessPassenger)
	visitEconomPassenger(ep *economPassenger)
}

// Passenger plane passenger
type Passenger interface {
	Accept(passengerVisitor)
}

type businessPassenger struct {
	Name string
}

// Accept visit concrete passenger
func (bp *businessPassenger) Accept(visitor passengerVisitor) {
	visitor.visitBusinessPassenger(bp)
}

type economPassenger struct {
	Name string
}

// Accept visit concrete passenger
func (this *economPassenger) Accept(visitor passengerVisitor) {
	visitor.visitEconomPassenger(this)
}

// PassengersList list of plane passengers
type PassengersList struct {
	passengers []Passenger
}

// Add append new passenger to list
func (pl *PassengersList) Add(pass Passenger) {
	pl.passengers = append(pl.passengers, pass)
}

// Accept visit all passengers by concrete visitor
func (pl *PassengersList) Accept(visitor passengerVisitor) {
	for _, pass := range pl.passengers {
		pass.Accept(visitor)
	}
}

// NewBusinessPassenger create new business passenger
func NewBusinessPassenger(name string) *businessPassenger {
	return &businessPassenger{Name: name}
}

// NewEconomPassenger create new business passenger
func NewEconomPassenger(name string) *economPassenger {
	return &economPassenger{Name: name}
}
