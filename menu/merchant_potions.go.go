package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

func buyRedPotion(p *player.Player) {
	if p.Money >= 20 {
		item := types.Item{Name: "Red potion", Type: "potion", BonusType: "poison", Bonus: 25}
		if p.Inventory.AddItem(item) {
			p.Money -= 20
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

func buyBag(p *player.Player) {
	if p.Money >= 40 {
		p.Money -= 40
		p.Inventory.MaxSize += 5
		fmt.Printf("Bag purchased! Inventory size : %d\n", p.Inventory.MaxSize)
	} else {
		fmt.Println("Not enough money.")
	}
}
