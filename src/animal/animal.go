package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	var name, action string

	fmt.Println("Enter an animal name (cow, bird, snake) followed by a space and an action (eat, move, speak).")
	fmt.Println("Examples: cow eat | bird move | snake speak")

	for {
		var animal Animal
		fmt.Printf("> ")
		fmt.Scanf("%s %s", &name, &action)

		switch name {
		case "cow":
			animal = Animal{food: "grass", locomotion: "walk", noise: "moo"}

		case "bird":
			animal = Animal{food: "worms", locomotion: "fly", noise: "peep"}
		case "snake":
			animal = Animal{food: "mice", locomotion: "slither", noise: "hsss"}
		default:
			fmt.Println("Invalid animal name")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "speak":
			animal.Speak()
		case "move":
			animal.Move()
		default:
			fmt.Println("Invalid action")
			continue
		}

	}
}
