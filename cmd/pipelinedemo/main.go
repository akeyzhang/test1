package main

import (
	"test1/pipeline"
	"fmt"
)
func MergeDemo() {
	p1 := pipeline.ArraySource(2, 10, 6, 3, 8, 21, 16, 7, 5)
	q1 := pipeline.InMemSort(p1)
	p2 := pipeline.ArraySource(4, 8, 1, 7, 3, 13, 22, 11, 12, 9)
	q2 := pipeline.InMemSort(p2)
	res := pipeline.Merge(q1, q2)
	for v := range res {
		fmt.Println(v)
	}
}

func main() {
	MergeDemo()

}

