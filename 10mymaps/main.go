package main

import "fmt"

func main() {
	// Creating a map using make
	studentAges := make(map[string]int)
	studentAges["Alice"] = 23
	studentAges["Bob"] = 20

	// Creating a map using a map literal
	studentGrades := map[string]int{
		"Alice": 95,
		"Bob":   85,
	}

	fmt.Println("Ages:", studentAges)
	fmt.Println("Grades:", studentGrades)

	fmt.Println("Bob's Grade:", studentGrades["Bob"])

	// Modifying an existing element
	studentGrades["Bob"] = 90

	// Adding a new element
	studentGrades["Charlie"] = 75

	fmt.Println("Updated Grades:", studentGrades)

	grade, exists := studentGrades["Bob"]
	if exists {
		fmt.Println("Bob's Grade:", grade)
	} else {
		fmt.Println("Bob not found")
	}

	delete(studentGrades, "Bob")

	fmt.Println("Updated Grades:", studentGrades)

	grad, exist := studentGrades["Bob"]
	if exist {
		fmt.Println("Bob's Grade:", grad)
	} else {
		fmt.Println("Bob not found")
	}

	// Loop through the map using a for range loop
	for name, grade := range studentGrades {
		fmt.Printf("%s's Grade: %d\n", name, grade)
	}
}
