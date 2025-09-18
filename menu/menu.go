package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

// Function to display the main menu
func ShowMenu(p *player.Player) {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Inventory")
		fmt.Println("2. Blacksmith")
		fmt.Println("3. Merchant")
		fmt.Println("4. Player statistics")
		fmt.Println("5. Resume the game")
		fmt.Println("6. Exit")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Display the player's inventory
			ShowPlayerInventory(&p.Inventory)
		case 2:
			// Blacksmith menu
			Blacksmith(p)
		case 3:
			// Merchant menu
			Merchant(p)
		case 4:
			// Display player information
			player.DisplayInfo(*p)
		case 5:
			// Resume the game
			return
		case 6:
			// Confirm exit
			if ConfirmExit() {
				return
			}
		default:
			fmt.Println("Invalid choice")
		}
	}
}
