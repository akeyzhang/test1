package main

import (
	"fmt"
	"time"
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

	t := time.Date(2018, 1, 12, 13, 5, 30, 0, time.Local)
	fmt.Println(t.Local())
}
