package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func paintNeeded(width float64, height float64) (area float64, err error) {
	//Func get 2 nubers and return area

	if width < 0 {
		return 0, fmt.Errorf("a width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %.2f is invalid", height)
	}
	area = width * height
	fmt.Println(area)
	return area / 10.0, nil
}

func checkError(err error) error {
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	fmt.Print("Input width: ")
	reader := bufio.NewReader(os.Stdin)
	width, err := reader.ReadString('\n')
	checkError(err)
	width = strings.TrimSpace(width)
	w, err := strconv.ParseFloat(width, 64)
	checkError(err)

	fmt.Print("Input height: ")
	reader = bufio.NewReader(os.Stdin)
	height, err := reader.ReadString('\n')
	checkError(err)
	height = strings.TrimSpace(height)
	h, err := strconv.ParseFloat(height, 64)

	amount, err := paintNeeded(w, h)
	checkError(err)

	fmt.Printf("For %.1f meters of wall, you need %d litters of paint\n", w*h, int(amount+1))
}

// func myFunc(i int, n int) (sum int) {
// 	//Func get 2 nubers and return square
// 	sum = i * n
// 	return sum
// }

// func main() {
// 	value := myFunc(5, 4)
// 	fmt.Println(value)
// }
