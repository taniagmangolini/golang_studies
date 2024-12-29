package main

import (
	"fmt"
)


func main() {
	var floatNumber float32
	
	fmt.Print("Enter a float number:")
	_, err := fmt.Scan(&floatNumber)
	if err != nil {
		fmt.Println("Invalid number!")
		return
	}
	fmt.Printf("Int number %d: ", int(floatNumber))
	fmt.Println()
	
}