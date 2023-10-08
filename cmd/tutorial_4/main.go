package main

import "fmt"

func main() {
	var intArr [3]int32
	intArr[0] = 123
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	var intArr2 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [3]int32{6, 7, 8}
	fmt.Println(intArr3)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])
	var age, ok = myMap2["Adam"]
	fmt.Println(age, ok)
	var age2, ok2 = myMap2["Jason"]
	fmt.Println(age2, ok2)

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
