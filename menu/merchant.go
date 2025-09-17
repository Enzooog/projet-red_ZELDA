package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

func Merchant(p *player.Player) {
	fmt.Println("\n--- MARCHAND FARAS ---")
	fmt.Printf("Bonjour ! Vos roupies : %d\n", p.Money)
	fmt.Println("1. Sheikah armor (50 roupies) - Stealth")
	fmt.Println("2. Stone armor (50 roupies) - +20 PV")
	fmt.Println("3. Bird armor (50 roupies) - +10 speed")
	fmt.Println("4. Green potion (15 roupies) - Treat 30 PV")
	fmt.Println("5. Red potion  (20 roupies) - Enemy poison")
	fmt.Println("6. bag (40 roupies) - +5 inventory slots")
	fmt.Println("7. back")
	fmt.Print("Choice : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		buySheikah(p)
	case 2:
		buyStone(p)
	case 3:
		buyBird(p)
	case 4:
		buyGreenPotion(p)
	case 5:
		buyRedPotion(p)
	case 6:
		buyBag(p)
	case 7:
		return
	default:
		fmt.Println("Invalid choice.")
	}
}

func buyGreenPotion(p *player.Player) {
	if p.Money >= 15 {
		item := types.Item{Name: "green potion", Type: "potion", BonusType: "heal", Bonus: 30}
		if p.Inventory.AddItem(item) {
			p.Money -= 15
		}
	} else {
		fmt.Println("Not enough money")
	}
}
