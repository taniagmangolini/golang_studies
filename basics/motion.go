package main

import (
	"fmt"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		finalPosition := s0 + v0 * t + (a * t * t)/2
		return finalPosition
	}
}


func main() {
	var acceleration float64
	var initialPosition  float64
	var initialVelocity  float64
	var t float64


	fmt.Print("Enter acceleration:")
	_, err := fmt.Scan(&acceleration)
	if err != nil {
		fmt.Println("Invalid the acceleration!")
		return
	}

	fmt.Print("Enter the initial velocity:")
	_, err = fmt.Scan(&initialVelocity)
	if err != nil {
		fmt.Println("Invalid velocity!")
		return
	}

	fmt.Print("Enter the initial position:")
	_, err = fmt.Scan(&initialPosition)
	if err != nil {
		fmt.Println("Invalid position!")
		return
	}

	fmt.Print("Enter the time:")
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Println("Invalid time!")
		return
	}

	fn := GenDisplaceFn(acceleration, initialVelocity, initialPosition)
	fmt.Println(fn(t))
	
}