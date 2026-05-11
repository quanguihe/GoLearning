package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitTime int) func(numList ...int) int {
	return func(numList ...int) int {
		time.Sleep(time.Duration(awaitTime) * time.Millisecond)
		var sum int
		for _, num := range numList {
			sum += num
		}
		return sum
	}
}

func main() {

	sum := awaitAdd(2)(1, 2, 3)
	fmt.Println(sum)

}
