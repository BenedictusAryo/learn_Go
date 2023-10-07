package main

import "fmt"

func main() {
	var intNum int16 = 327
	intNum += 1
	fmt.Println(intNum)

	var floatNum float32 = 3.14
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)
	fmt.Println(float32(intNum1) / float32(intNum2))

	var myString string = `Hello, 
	World!`
	fmt.Println(myString)

	var intNum3 int
	fmt.Println(intNum3)

	var strvar = "text"
	fmt.Println(strvar)

	strvar2 := "text2"
	fmt.Println(strvar2)

	var1, var2 := 7, 8
	fmt.Println(var1, var2)
}
