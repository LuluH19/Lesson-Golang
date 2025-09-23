package main

import (
	"fmt"
)

// Affiche le carburant restant
func fuelGauge(fuel int) {
	fmt.Println("Fuel remaining:", fuel)
}

// Retourne le carburant nécessaire selon la planète
func calculateFuel(planet string) int {
	fuel := 0
	switch planet {
	case "Venus", "Mercury", "Mars":
		// as before
	default:
		fmt.Println("Unknown planet. No fuel calculated.")
		fuel = 0
	}
	return fuel
}

// Message d’accueil à destination
func greetPlanet(planet string) {
	fmt.Println("Welcome to", planet, "— enjoy your stay among the stars!")
}

// Message d’impossibilité de vol
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Coordonne le vol et renvoie le carburant restant
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
	// Carburant initial
	fuel := 1000000

	// Choix de la planète
	planetChoice := "Venus"

	// Vol vers la planète (met à jour le carburant)
	fuel = flyToPlanet(planetChoice, fuel)

	// Vérifie la jauge de carburant
	fuelGauge(fuel)

	// Bonus
	planetChoice = "Mars"
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
