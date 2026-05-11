package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func add(numberList ...int) int {
	var sum int
	for _, number := range numberList {
		sum += number
	}
	/*	for i := range numberList {
		sum += numberList[i]
	}*/
	return sum
}

func main() {
	sayHello()

	fmt.Println(add(1, 2, 3))

	//匿名函数
	var add = func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))

}
