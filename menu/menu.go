package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func ShowMenu(p *player.Player) {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Inventaire")
		fmt.Println("2. Forgeron")
		fmt.Println("3. Marchand")
		fmt.Println("4. Stats du joueur")
		fmt.Println("5. Reprendre le jeu")
		fmt.Println("6. Quitter")
		fmt.Print("Choix : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			p.Inventory.ShowInventory()
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
			fmt.Println("Choix invalide.")
		}
	}
}
