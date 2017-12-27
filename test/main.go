package main

import (
	"fmt"
	"time"
)

func main() {
	a:=[]int{1,2,3,4,5,6,7,8,9}
	b:=[]int{0,0,0,0,0,0,0,0,0}

	go func() {
		for i,v:=range a{
			b[i]=v
		}
	}()
	time.Sleep(time.Millisecond)

	for i,v:=range b{
		fmt.Println(i,v)
	}
}
