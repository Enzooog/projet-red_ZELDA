package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

func buySheikah(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Sheikah armor", Type: "armor", BonusType: "stealth", Bonus: 1}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

func buyStone(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Stone Armor", Type: "armor", BonusType: "pv", Bonus: 20}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.PvMax += 20
			p.Pv += 20
			fmt.Println("PV max increased !")
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

func buyBird(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Bird Armor", Type: "armor", BonusType: "speed", Bonus: 10}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.Speed += 10
			fmt.Println("Increased speed !")
		}
	} else {
		fmt.Println("Not enough money.")
	}
}
