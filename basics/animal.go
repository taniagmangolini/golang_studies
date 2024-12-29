package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "warms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	validOptions := []string{"eat", "move", "speak"}
	validAnimals := []string{"cow", "bird", "snake"}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		request, _ := reader.ReadString('\n') 
		request = strings.TrimSpace(request) 
		requestParts := strings.Split(request, " ")
		if len(requestParts) != 2 {
			fmt.Println("Invalid number of arguments! Your request should be in the format: animal option")
			return
		}
		isValidRequest := contains(validAnimals, requestParts[0]) && contains(validOptions, requestParts[1])
		if !isValidRequest {
			fmt.Println("Invalid arguments! Valid animals ", validAnimals, ". Valid options: ", validOptions)
			return
		}
		var animal Animal
		switch requestParts[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Println("Invalid animal!")
		}

		switch requestParts[1] {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			fmt.Println("Invalid option!")
		}
	}

}

func contains(s []string, val string) bool {
    for _, a := range s {
        if a == val {
            return true
        }
    }
    return false
}