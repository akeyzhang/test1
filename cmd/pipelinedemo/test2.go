package main

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

type aa int

func (a aa) test() {
	fmt.Println("a is :", a)
}

func main() {
	var a aa = 3
	a.test()
}
