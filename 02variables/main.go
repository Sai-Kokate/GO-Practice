package main

import "fmt"

func main() {
	var name string = "Saibalaji"
	fmt.Println(name)
	fmt.Printf("The type of the variable is: %T \n", name)

	var age int = 18
	fmt.Println(age)
	fmt.Printf("The type of the variable is: %T \n", age)

	var floatAge float64 = 23423.32434323424332
	fmt.Println(floatAge)
	fmt.Printf("The type of the variable is: %T \n", floatAge)

	var isCoder bool = true
	fmt.Println(isCoder)
	fmt.Printf("The type of the variable is: %T \n", isCoder)

	// default type
	var str string
	fmt.Println(str)
	fmt.Printf("The type of the variable is: %T \n", str)

	fristName, isAdult, newAge := "sai", true, 22
	fmt.Println(fristName)
	fmt.Printf("The type of the variable is: %T \n", fristName)
	fmt.Println(isAdult)
	fmt.Printf("The type of the variable is: %T \n", isAdult)
	fmt.Println(newAge)
	fmt.Printf("The type of the variable is: %T \n", newAge)

}
