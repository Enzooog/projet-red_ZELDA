package menu

import (
	"fmt"
	"os"
)

// Function to confirm exiting the game
func ConfirmExit() bool {
	fmt.Println("\nAre you sure you want to leave?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Print("Choice: ")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		// Exit the game
		fmt.Println("Goodbye and thank you for playing!")
		os.Exit(0)
		return true
	case 2:
		// Stay in the game
		return false
	default:
		fmt.Println("Invalid choice.")
	}
	return false
}
