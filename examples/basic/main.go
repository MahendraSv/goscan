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
}
