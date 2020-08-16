package main

import "fmt"

func main() {
	// using var
	var name = "Benedict"
	var age int64 = 26

	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}
