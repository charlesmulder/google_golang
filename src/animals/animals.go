package main

import (
	"fmt"
)

type Animal interface {
	Eat() 
	Move()
	Speak()
}

// Cow
type Cow struct {}

func (c Cow ) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird
type Bird struct {}

func (b Bird ) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake
type Snake struct {}

func (s Snake ) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func getAnimal(a string) (Animal, error) {
	switch a {
	case "cow":
		return Cow{}, nil
	case "bird":
		return Bird{}, nil
	case "snake":
		return Snake{}, nil
	default: 
		return nil, fmt.Errorf("Unknown animal: %s", a)
	}
}

func callMethod(a Animal, m string) error {
	switch m {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default: 
		return fmt.Errorf("Unknown method: %s", m)
	}
	return nil
}

func main() {

	animals := map[string]Animal{}

	var command, name, arg string 

	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s %s", &command, &name, &arg)

		switch command {

		case "newanimal":
			animal, err := getAnimal(arg)
			if err != nil  {
				fmt.Println(err)
				return
			}
			animals[name] = animal
			fmt.Println("Created it!")

		case "query":
			if err := callMethod(animals[name], arg ); err != nil {
				fmt.Println(err)
			}

		default:
			fmt.Println("Unknown command. Valid commands are 'newanimal' or 'query'.")
		}
	}

}
