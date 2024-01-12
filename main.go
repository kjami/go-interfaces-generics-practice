package main

import (
	"fmt"

	"example.com/interfaces-generics-practice/generics"
	"example.com/interfaces-generics-practice/vehicle"
)

func displayVehicle(v vehicle.Vehicle) {
	v.Display()
}

func main() {
	teslaM3 := vehicle.NewCar("Tesla Model 3", 2023, 258)
	crv2023 := vehicle.NewCar("Honda CR-V", 2023, 190)
	randomBus := vehicle.NewBus("Random bus", 2007, 500, 100)

	displayVehicle(teslaM3)
	displayVehicle(crv2023)
	displayVehicle(randomBus)

	fmt.Println(generics.Add(1, 2))
	fmt.Println(generics.Add(1.4, 2.2))
	fmt.Println(generics.Add("This is a", " string concat"))
	fmt.Println(generics.Square(9))
	fmt.Println(generics.Square(1.2))
	fmt.Println(generics.Square("random"))
}
