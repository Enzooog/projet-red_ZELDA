package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func ShowMenu(p *player.Player) {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Inventory")
		fmt.Println("2. Blacksmith")
		fmt.Println("3. merchant")
		fmt.Println("4. player statistics")
		fmt.Println("5. resume the game ")
		fmt.Println("6. exit")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ShowPlayerInventory(&p.Inventory)
		case 2:
			Blacksmith(p)
		case 3:
			Merchant(p)
		case 4:
			player.DisplayInfo(*p)
		case 5:
			return
		case 6:
			if ConfirmExit() {
				return
			}
		default:
			fmt.Println("invalid choice")
		}
	}
}
