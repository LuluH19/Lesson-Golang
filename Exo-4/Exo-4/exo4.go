package main

import (
	"fmt"
)

func fuelGauge(fuel int) {
	fmt.Println("Fuel remaining:", fuel)
}

func calculateFuel(planet string) int {
	fuel := 0
	switch planet {
	case "Venus", "Mercury", "Mars":

	default:
		fmt.Println("Unknown planet. No fuel calculated.")
		fuel = 0
	}
	return fuel
}

func greetPlanet(planet string) {
	fmt.Println("Welcome to", planet, "— enjoy your stay among the stars!")
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining int
	var fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {

	fuel := 1000000

	planetChoice := "Venus"

	fuel = flyToPlanet(planetChoice, fuel)

	fuelGauge(fuel)

	planetChoice = "Mars"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
