package main

import (
	"fmt"
)

type Comparable[T any] interface {
	Compare(T) bool
}

func main() {
	xs := Ints{1, 2, 3}
	ys := Ints{1, 2, 3}
	fmt.Println(xs.Compare(ys))
}

