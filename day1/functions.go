package main

import (
	"fmt"
)

func myFunc(i int, n int) (sum int) {
	sum = i * n
	return sum
}

func main() {
	value := myFunc(5, 4)
	fmt.Println(value)
}
