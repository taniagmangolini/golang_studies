package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct { 
	Name string
}

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct { 
	Name string
}

func (a Bird) Eat() {
	fmt.Println("warms")
}

func (a Bird) Move() {
	fmt.Println("fly")
}

func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct { 
	Name string
}

func (a Snake) Eat() {
	fmt.Println("mice")
}

func (a Snake) Move() {
	fmt.Println("slither")
}

func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	animals := make([]Animal, 0)
	for {
		fmt.Print("> ")
		request, _ := reader.ReadString('\n') 
		request = strings.TrimSpace(request) 
		requestParts := strings.Split(request, " ")
		if len(requestParts) != 3 {
			fmt.Println("Invalid number of arguments!")
			return
		}
		switch requestParts[0] {
		case "newanimal":
			animal := CreateNewAnimal(requestParts[1], requestParts[2])
			if animal != nil {
				animals = append(animals, animal)
				fmt.Println("Created it!")
			}
		case "query":
			foundAnimals := FindAnimal(animals, requestParts[1])
			if len(foundAnimals) == 0 {
				fmt.Println("Animals not found")
			}
			QueryAnimal(foundAnimals, requestParts[2])
		default:
			fmt.Println("Invalid request!")
		}
	}
}

func CreateNewAnimal(name string, option string) Animal {
	switch option {
	case "cow":
		return Cow{Name: name}
	case "bird":
		return Bird{Name: name}
	case "snake":
		return Snake{Name: name}
	default:
		fmt.Println("Invalid animal type!")
		return nil
	}
}

func FindAnimal(animals []Animal, name string) []Animal {
	found := []Animal{}
    for _, animal := range animals {
		var animalName string
		switch a := animal.(type) {
		case Cow:
			animalName = a.Name
		case Bird:
			animalName = a.Name
		case Snake:
			animalName = a.Name
		default:
		}
		if animalName == name {
			found = append(found, animal)
		}
    }
    return found
}

func QueryAnimal(animals []Animal, option string) {
	for _, animal := range animals {
		switch option {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid action for this animal!")
		}
	}
}
