package main

import "fmt"

func main() {
	fmt.Println("Poniters!")

	var ptr *int // pointer ptr pointing to an address that stores int value

	fmt.Println("ptr inital value: ", ptr)

	myNum := 22

	var newptr = &myNum

	fmt.Println("Address newptr: ", newptr)
	fmt.Println("Value *newptr: ", *newptr)

	*newptr = *newptr + 8
	fmt.Println("Updated mynum: ", myNum)

}
