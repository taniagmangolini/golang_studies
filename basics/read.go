package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Fname string
	Lname string
}

func main() {
	const maxNameLength = 20
	persons := make([]Person, 0)
	
	// asks for the filename
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the filename: ")
	scanner.Scan()
	filename := scanner.Text()
	if filename == "" {
		fmt.Println("Invalid filename ")
		return
	}

	// read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// iterate over the file
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			fmt.Println("Error parsing the file. Invalid line.")
			return
		}
		if len(parts[0]) > 20 ||  len(parts[1]) > 20{
			fmt.Println("Error parsing the file. The maximum Fname/Lane size is 20.")
			return
		}
		persons = append(persons, Person{Fname: parts[0], Lname: parts[1]})
	}

	// Print the Persons's slice
	for _, person := range persons {
		fmt.Println(person.Fname, person.Lname)
	}

}