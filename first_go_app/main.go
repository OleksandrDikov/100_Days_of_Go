package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	fmt.Println(year, "-", month, "-", day)

	// fmt.Println("Hello World!")
	// fmt.Println("Second Line.")
	// fmt.Println(math.Floor(5.15))
	// fmt.Println(strings.Title("go programing lang"))
	// fmt.Println(reflect.TypeOf(5))
	// fmt.Println(reflect.TypeOf(5.15))
	// fmt.Println(reflect.TypeOf("Text"))
	// fmt.Println(reflect.TypeOf(true))
}
