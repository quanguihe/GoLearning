package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) SetAge(age int) {
	s.Age = age
}
func (s *Student) SetAge1(age int) {
	s.Age = age
}

func main() {
	s := Student{
		Name: "枫枫",
		Age:  21,
	}
	s.SetAge(18)
	fmt.Println(s.Age)
	s.SetAge1(18)
	fmt.Println(s.Age)
}
