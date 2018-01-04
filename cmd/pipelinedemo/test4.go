package main

import (
	"strconv"
	"fmt"
)

type testfunc func(int) string

func test2(a int) string {
	return "[test2] you input :"+strconv.Itoa(a)
}

func test1(a int) string {
	return "[test1] you input :"+strconv.Itoa(a)
}



func say(f testfunc, a int) {
	fmt.Println(f(a))
}

func main() {
  //say(test2,4)
  f:=test1
  str:=f(3)
  fmt.Println(str)

}
