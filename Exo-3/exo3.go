package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	isHeistOn := true

	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println(" You slipped past the guards. First hurdle cleared!")
	} else {
		isHeistOn = false
		fmt.Println(" Caught by the guards… Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println(" Vault opened! Grab and GO!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println(" The vault won't budge… aborting the mission.")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println(" Silent alarm tripped at the exit!")
		case 1:
			isHeistOn = false
			fmt.Println(" Police blockade outside—no way through.")
		case 2:
			isHeistOn = false
			fmt.Println(" Cameras caught your plates. Not good.")
		case 3:
			isHeistOn = false
			fmt.Println(" Extra guards showed up at the last second.")
		default:
			fmt.Println(" Start the getaway car! You’re out!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("Take counted: $%d\n", amtStolen)
	}

	fmt.Println("isHeistOn is currently:", isHeistOn)
}
