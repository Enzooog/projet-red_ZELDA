package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

// Buy a red potion
func buyRedPotion(p *player.Player) {
	if p.Money >= 20 {
		// Create the item
		item := types.Item{Name: "Red potion", Type: "potion", BonusType: "poison", Bonus: 25}
		if p.Inventory.AddItem(item) {
			p.Money -= 20 // Deduct money
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

func buyGreenPotion(p *player.Player) {
	if p.Money >= 15 {
		item := types.Item{Name: "Green potion", Type: "potion", BonusType: "heal", Bonus: 30}
		if p.Inventory.AddItem(item) {
			p.Money -= 15
		}
	} else {
		fmt.Println("Not enough money.")
	}
}
