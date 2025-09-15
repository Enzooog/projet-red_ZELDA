package menu

import "fmt"

func Show() {
	for {
		fmt.Println(" MENU ")
		fmt.Println("1. Inventory")
		fmt.Println("2. Blacksmith")
		fmt.Println("3. Player Information")
		fmt.Println("4. Resume the Game")
		fmt.Println("5. Exit the Game")

		var key int
		fmt.Print("Choice: ")
		fmt.Scan(&key)

		switch key {
		case 1:
			fmt.Println("Opening inventory")
		case 2:
			fmt.Println("Going to the blacksmith")
		case 3:
			merchant()
		case 4:
			fmt.Println("Showing player information")
		case 5:
			fmt.Println("Resuming the game")
		case 6:
			confirmExit()
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
