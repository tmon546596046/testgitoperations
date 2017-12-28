package model

import (
	"fmt"
)

type Animal struct {
	Name        string
	Description string
}

func (a *Animal) Cry() {
	fmt.Println("miao miao")
}
