package main

import "fmt"

func main() {
	var x [3][5]int
	fmt.Println(x)
	x[1][2] = 25
	fmt.Println(x[1][2])
}
