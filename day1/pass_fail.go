package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	if grade == 100 {
		fmt.Println("Perfect!")
	} else if grade >= 60 {
		fmt.Println("You pass!")
	} else {
		fmt.Println("You fail!")
	}
}
