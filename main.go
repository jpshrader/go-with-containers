package main

import (
	"fmt"
)

func isEven(x int) bool {
	return x % 2 == 0
}

func isOdd(x int) bool {
	return !isEven(x)
}

func main() {
	numbers := Slice[int]{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenNumbers := numbers.Where(isEven)
	oddNumbers := numbers.Where(isOdd)

	fmt.Println(fmt.Sprintf("even numbers: %v", evenNumbers))
	fmt.Println(fmt.Sprintf("odd numbers:  %v", oddNumbers))
}

type Slice[T any] []T

type WhereFilter[T any] func(x T) bool

func (slice Slice[T]) Where(filter WhereFilter[T]) (ret []T) {
	for _, item := range slice {
		if filter(item) {
			ret = append(ret, item)
		}
	}
	return
}