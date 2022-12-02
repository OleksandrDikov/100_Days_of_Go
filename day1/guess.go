package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var success bool
	t := time.Now().Unix()
	rand.Seed(t)
	target := rand.Intn(100) + 1
	fmt.Print("Try to guess number from 1 to 100: ")

	for x := 10; x > 0; x-- {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		userNum, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if userNum == target {
			fmt.Println("Good gob! You guessed it!")
			success = true
			break
		} else {
			if userNum > target {
				fmt.Println("You guess is to High!")
			} else {
				fmt.Println("You guess is to Low!")
			}
			reader = bufio.NewReader(os.Stdin)
		}

		fmt.Println("You have", x-1, "attempts.")
		fmt.Print("Make a guess: ")
	}
	if !success {
		fmt.Println("Sorry, you didn't guess the number.")
	}
	fmt.Println("The number was", target)
}
