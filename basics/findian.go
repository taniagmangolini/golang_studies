package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid string!")
		return
	}

	input = strings.TrimSpace(input)

	inputLowerCase := strings.ToLower(input)

	if strings.HasPrefix(inputLowerCase, "i") && strings.Contains(inputLowerCase, "a") && strings.HasSuffix(inputLowerCase, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
	fmt.Println()
}
