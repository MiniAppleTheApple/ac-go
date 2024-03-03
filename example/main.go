package main

import (
	"fmt"
)

type Comparable[T any] interface {
	Compare(T) bool
}

type Ints []int

func (ints Ints) Compare(another Ints) bool {
	if len(ints) != len(another) {
		return false
	}

	for index := range ints {
		if ints[index] != another[index] {
			return false
		}
	}

	return true
}

func main() {
	xs := Ints{1, 2, 3}
	ys := Ints{1, 2, 3}
	fmt.Println(xs.Compare(ys))
}

