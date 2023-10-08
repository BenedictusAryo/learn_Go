package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var name string
	var age int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(strings.TrimSpace(name), "\n")
	for {
		fmt.Println("Enter your age:")
		_, err := fmt.Scan(&age)
		if err == nil {
			break
		}
		fmt.Println("Please enter a valid age.")
	}
	birthYear := time.Now().Year() - age
	fmt.Printf("Hello %v, you were born in %v.\n", strings.Title(name), birthYear)
}
