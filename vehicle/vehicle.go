package vehicle

import "fmt"

type Vehicle interface {
	Start() string
	Honk() string
	Display()
}

type CommericialVehicle interface {
	Load(int) string
	Vehicle
}

type Car struct {
	model        string
	releasedYear int
	horsePower   int
}

func NewCar(model string, releasedYear int, horsePower int) *Car {
	return &Car{
		model,
		releasedYear,
		horsePower,
	}
}

func (c Car) Start() string {
	return "Car: started"
}

func (c Car) Honk() string {
	return "Car honk"
}

func (c Car) Display() {
	fmt.Println(c.model, c.releasedYear, c.horsePower)
}

type Bus struct {
	model           string
	releasedYear    int
	horsePower      int
	seatingCapacity int
}

func NewBus(model string, releasedYear int, horsePower int, seatingCapacity int) *Bus {
	return &Bus{
		model,
		releasedYear,
		horsePower,
		seatingCapacity,
	}
}

func (b Bus) Start() string {
	return "Bus: started"
}

func (b Bus) Honk() string {
	return "Bus honk"
}

func (b Bus) Load(people int) string {
	return "Loaded: " + fmt.Sprint(people) + " out of " + fmt.Sprint(b.seatingCapacity)
}

func (b Bus) Display() {
	fmt.Println(b.model, b.releasedYear, b.horsePower, b.seatingCapacity)
}
