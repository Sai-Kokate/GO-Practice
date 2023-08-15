package main

import "fmt"

func main() {
	fmt.Println("Loops in Go!")

	months := []string{"Jan", "Feb", "March", "Apr", "May", "June"}

	// for loop in go
	for m := 0; m < len(months); m++ {
		fmt.Println(months[m])
	}

	// for range loops

	for index, val := range months {
		fmt.Printf("Index: %v , Month: %v \n", index, val)
	}

	for index := range months {
		fmt.Println(months[index])

	}

	for _, val := range months {
		fmt.Printf("Month: %v \n", val)
	}

	// while loop
	count := 0
	for count < 5 {

		if count == 1 {
			count++
			continue
		}

		if count == 3 {
			goto labl
		}
		fmt.Println("count: ", count)
		count++
	}

labl:
	fmt.Println("Jumped to the lable using goto")

}
