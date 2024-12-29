package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"sort"
)

func main() {
	numbers := make([]int, 0, 3)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a integer number (type 'X' to exit): ")
		input, _ := reader.ReadString('\n') 
		input = strings.TrimSpace(input) 

		if strings.ToLower(input) == "x" { // Check if input is 'x'
			break
		}
		fmt.Println(input)

		number, err := strconv.Atoi(input)
		if err == nil {
			numbers = append(numbers, number)
			sort.Ints(numbers)
			fmt.Println(numbers) 
		} else {
			fmt.Println("The input is not an integer. Exiting...")
			break
		}
	}
}
