package main

import "fmt"

func main() {
	fmt.Println("Arrays!")

	var strarr [4]string

	strarr[0] = "Sai"
	strarr[1] = "Balaji"
	strarr[2] = "Kokate"

	fmt.Println("Array : ", strarr)
	fmt.Println("Array length : ", len(strarr))

	var newarr = [4]string{"Hi", "Hello", "Hola"}
	fmt.Println("Newarr: ", newarr)
	fmt.Println("New Array length : ", len(newarr))
}
