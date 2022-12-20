package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func setNumber() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	return number
}

func main() {
	fmt.Print("Enter first number: ")
	numb1 := setNumber()
	fmt.Print("Enter second number: ")
	numb2 := setNumber()
	sum := numb1 + numb2

	fmt.Println("Sum of two numbers = ", sum)
}
