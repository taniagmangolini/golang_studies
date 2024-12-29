package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func Swap(numbers []int, index int) {
	numbers[index], numbers[index + 1] = numbers[index + 1], numbers[index]
}

func BubbleSort(numbers []int) {
	if len(numbers) <= 1 {
		return 
	}
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers) - i - 1; j++ {
			if numbers[j] > numbers[j + 1] {
				Swap(numbers, j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numbers := make([]int, 0)
	fmt.Println("Enter integers numbers for sorting (Type 'X' to finish. There is a limit of 10 numbers): ")
	for {
		
		input, _ := reader.ReadString('\n') 
		input = strings.TrimSpace(input) 

		if strings.ToLower(input) == "x" { // Check if input is 'x'
			break
		}
		number, err := strconv.Atoi(input)
		if err == nil {
			numbers = append(numbers, number)
		} else {
			fmt.Println("The input is not an integer. Exiting...")
			break
		}
		if len(numbers) == 10 { 
		break
	}
	}
	BubbleSort(numbers)
	fmt.Println(numbers) 
}