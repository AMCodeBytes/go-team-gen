package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int

	fmt.Println("Welcome to Team Generator")
	fmt.Println("1. List")
	fmt.Println("2. Create New")
	fmt.Println("3. Add")
	fmt.Println("4. Remove")
	fmt.Println("5. Generate Teams")
	fmt.Println("6. Quit")
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("You have chosen option 1.")
	case 2:
		fmt.Println("You have chosen option 2.")
	case 3:
		fmt.Println("You have chosen option 3.")
	case 4:
		fmt.Println("You have chosen option 4.")
	case 5:
		fmt.Println("You have chosen option 5.")
	case 6:
		fmt.Println("Goodbye, see you soon.")
		os.Exit(0)
	default:
		panic("Not a valid input")
	}
}
