package main

import "fmt"

type SingInterface interface {
	Sing()
}

type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name + "在唱歌")
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name + "在唱歌")
}

// 空接口
func Print(c interface{}) {
	fmt.Printf("%+v\n", c)
}

func main() {
	s1 := SingInterface(Chicken{Name: "chicken"})
	s1.Sing()

	s2 := SingInterface(Cat{Name: "cat"})
	s2.Sing()

	//空接口
	c := Chicken{Name: "chicken"}
	Print(c)
}
