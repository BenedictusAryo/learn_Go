package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello, World!"
	printMe(printValue)

	numerator, denominator := 11, 0
	var result, remainder, err = intDivision(numerator, denominator)
	// using if else
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	// using switch
	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	// using switch remainder condition
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close ")
	default:
		fmt.Printf("The division was not exact")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
