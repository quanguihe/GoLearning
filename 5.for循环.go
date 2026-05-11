package main

import "fmt"

func main() {

	var sum = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	s := []string{"枫枫", "知道"}
	for index, s2 := range s {
		fmt.Println(index, s2)
	}
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
}
