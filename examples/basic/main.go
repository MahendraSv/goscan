package main

import (
	"fmt"

	"github.com/MahendraSv/goscan/scanner"
)

func main() {
	in := scanner.New()

	name := in.String("Enter name: ")
	age := in.Int("Enter age: ")
	isDev := in.Bool("Are you a developer? ")

	fmt.Println("\n--- Result ---")
	fmt.Println(name, age, isDev)

	// Using generics
	name1 := scanner.Read[string](in, "Name: ")
	age1 := scanner.Read[int](in, "Age: ")
	height1 := scanner.Read[float64](in, "Height: ")
	isDev1 := scanner.Read[bool](in, "Developer? ")

	fmt.Println("\n--- Result ---")
	fmt.Println(name1, age1, height1, isDev1)
}
