package main

import (
	"fmt"
)

func main() {
	var name string
	var age int

	// Demander le nom
	fmt.Println("What is your name?")
	fmt.Scan(&name)

	fmt.Printf("Hello %s!\n", name)

	// Demander l'âge
	fmt.Println("What is your age?")
	fmt.Scan(&age)

	// Afficher avec Printf
	fmt.Printf("%s is %d years old.\n", name, age)

	// Reformater avec Sprintf
	newMessage := fmt.Sprintf("User info → Name: %s | Age: %d", name, age)
	fmt.Println(newMessage)
}
