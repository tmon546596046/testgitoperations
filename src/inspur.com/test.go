package main

import (
	"fmt"
	"inspur.com/model"
)

func main() {
	fmt.Println("hello world 1")
	fmt.Println("hello world 2")
	a := new(model.Animal)
	a.Name = "kitty"
	fmt.Printf("hello %s\n", a.Name)
	a.Cry()
}
