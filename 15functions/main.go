package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	sumresult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sumresult)

	newresult, success := divide(10, 0)
	if success {
		fmt.Println("Result:", newresult)
	} else {
		fmt.Println("Division by zero.")
	}

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
