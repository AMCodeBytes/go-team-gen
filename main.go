package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	var players []string

	for {
		fmt.Println("Welcome to Team Generator")
		fmt.Println("1. List")
		fmt.Println("2. Create New")
		fmt.Println("3. Add")
		fmt.Println("4. Remove")
		fmt.Println("5. Remove All")
		fmt.Println("6. Generate Teams")
		fmt.Println("7. Quit")
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("-----  List Players  -----")
			fmt.Println(players)
		case 2:
			fmt.Println("----- Create New Players  -----")
			clearPlayers(&players)
			newPlayers(&players)
		case 3:
			fmt.Println("----- Add Players  -----")
			newPlayers(&players)
		case 4:
			fmt.Println("You have chosen option 4.")
		case 5:
			fmt.Println("----- Remove All Players  -----")
			clearPlayers(&players)
		case 6:
			fmt.Println("You have chosen option 5.")
		case 7:
			fmt.Println("Goodbye, see you soon.")
			os.Exit(0)
			return
		default:
			panic("Not a valid input")
		}
	}
}

func newPlayers(list *[]string) {
	var input string

	fmt.Println("enter \"q\" to quit")

	for {
		fmt.Print("New player: ")
		fmt.Scan(&input)

		if input == "q" {
			return
		}

		*list = append(*list, input)
	}
}

func clearPlayers(list *[]string) {
	*list = nil
}
