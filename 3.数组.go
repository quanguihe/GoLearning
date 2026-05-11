package main

import "fmt"

func main() {

	array := [3]int{1, 2, 3}
	fmt.Println(array)
	slice := array[:]
	fmt.Println(slice)

	var nameList = make([]string, 3)
	fmt.Println(nameList)

	ageList := make([]int, 3)
	fmt.Println(ageList)
}
