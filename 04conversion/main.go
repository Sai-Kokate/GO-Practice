package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the course rating system!")
	fmt.Println("Please the course between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter you rating : ")

	rating, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
