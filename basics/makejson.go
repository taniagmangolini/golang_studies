package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
	"strings"

)

func main() {
	questions := map[string]string {
		"name": "Enter the name:", 
		"address": "Enter the address:",
	}
	personalInfo := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	for field, question := range questions {
		fmt.Print(question)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Errror scanning " + field)
			return
		}
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Invalid " + field)
			return
		}
		personalInfo[field] = input
	}

	personalInfoJson, err := json.Marshal(personalInfo)
	if err != nil {
		fmt.Println("Error converting to json.")
		return
	}

	fmt.Print(string(personalInfoJson))
	fmt.Println()
	
}



/*
# Alternative
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	person := make(map[string]string)

	fmt.Print("Enter a name: ")
	scanner.Scan()
	person["name"] = scanner.Text()

	fmt.Print("Enter an address: ")
	scanner.Scan()
	person["address"] = scanner.Text()

	m, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(m))
}

*/