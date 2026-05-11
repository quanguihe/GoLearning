package main

import "fmt"

func main() {
	fmt.Println("请输入你的年龄：")
	var age int
	fmt.Scan(&age)

	switch {
	case age <= 0:
		fmt.Println("未出生")
		fallthrough
	case age <= 18:
		fmt.Println("未成年")
		fallthrough
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")
	}
}
