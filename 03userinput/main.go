package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter you Name: ")

	reader := bufio.NewReader(os.Stdin)

	inp, _ := reader.ReadString('\n')

	fmt.Println("You Entered: ", inp)
	fmt.Printf("Type of inp: %T", inp) // even if you enter a int, the type comes as string
}
