package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

func Merchant(p *player.Player) {
	fmt.Println("\n--- MARCHAND FARAS ---")
	fmt.Printf("Bonjour ! Vos roupies : %d\n", p.Money)
	fmt.Println("1. Armure Sheikah (50 roupies) - Furtivité")
	fmt.Println("2. Armure Pierre (50 roupies) - +20 PV")
	fmt.Println("3. Armure Oiseau (50 roupies) - +10 Vitesse")
	fmt.Println("4. Potion verte (15 roupies) - Soigne 30 PV")
	fmt.Println("5. Potion rouge (20 roupies) - Poison ennemi")
	fmt.Println("6. Sac (40 roupies) - +5 slots inventaire")
	fmt.Println("7. Retour")
	fmt.Print("Choix : ")

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
		fmt.Println("Choix invalide.")
	}
}

func buyGreenPotion(p *player.Player) {
	if p.Money >= 15 {
		item := types.Item{Name: "Potion verte", Type: "potion", BonusType: "heal", Bonus: 30}
		if p.Inventory.AddItem(item) {
			p.Money -= 15
		}
	} else {
		fmt.Println("Pas assez d'argent.")
	}
}
