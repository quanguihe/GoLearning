package main

import "fmt"

func change(name string) {
	fmt.Println("%p", &name)
	fmt.Println(name)
}

func main() {
	name := "zhangsan"
	change(name)
	fmt.Println("%p", &name)
}
