package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	sai := User{"Sai", "sai@go.dev", true, 22}
	fmt.Println(sai)
	fmt.Printf("Sai details are: %+v\n", sai)
	fmt.Printf("Name is %v and email is %v.", sai.Name, sai.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
