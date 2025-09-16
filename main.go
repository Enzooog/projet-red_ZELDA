package main

import (
	"fmt"
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {
	p := player.CreatePlayer()
	player.DisplayInfo(p)

	p.Pv = 0
	player.IsDead(&p)
	player.TakePot(&p)
	for {
		fmt.Println("\nQue veux-tu faire ?")
		fmt.Println("1. Afficher infos joueur")
		fmt.Println("2. Acheter un objet ")
		fmt.Println("3. Afficher l'inventaire et équiper un objet")
		fmt.Println("4. Quitter")
		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			player.DisplayInfo(p)
		case 2:
			menu.BuyStone(&p.Money, &p.Inventory)
		case 3:
			equipped := p.Inventory.EquipItem()
			if equipped != nil && equipped.BonusType == "pv" {
				p.PvMax += equipped.Bonus
				p.Pv += equipped.Bonus
				fmt.Printf("Bonus de PV appliqué : %d ! Nouveaux PV max : %d\n", equipped.Bonus, p.PvMax)
			}
		case 4:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
