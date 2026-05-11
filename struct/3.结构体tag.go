package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}

func main() {
	user := User{Name: "zhangsan", Age: 18, Password: "123456"}
	bytes, _ := json.Marshal(user)
	fmt.Println(string(bytes))
}
