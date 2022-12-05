package main

import (
	"fmt"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evenNumbers := Where(isEven, items)

	fmt.Printf("even numbers: %v", evenNumbers)
}

func isEven(x int) bool {
	return x % 2 == 0
}

type WhereFilter[T any] func(x T) bool

func Where[T any](filter WhereFilter[T], source []T,) ([]T) {
	returnValues := make([]T, 0)

	for _, item := range source {
		if filter(item) {
			returnValues = append(returnValues, item)
		}
	}

	return returnValues
}