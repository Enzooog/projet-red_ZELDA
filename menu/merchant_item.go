package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

// Buy the Sheikah armor
func buySheikah(p *player.Player) {
	if p.Money >= 50 {
		// Create the item
		item := types.Item{Name: "Sheikah Armor", Type: "armor", BonusType: "stealth", Bonus: 1}
		if p.Inventory.AddItem(item) {
			p.Money -= 50 // Deduct money
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

// Buy the Stone armor
func buyStone(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Stone Armor", Type: "armor", BonusType: "pv", Bonus: 20}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.PvMax += 20 // Increase max HP
			p.Pv += 20    // Restore current HP
			fmt.Println("Max HP increased!")
		}
	} else {
		fmt.Println("Not enough money.")
	}
}

// Buy the Bird armor
func buyBird(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Bird Armor", Type: "armor", BonusType: "speed", Bonus: 10}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.Speed += 10 // Increase speed
			fmt.Println("Speed increased!")
		}
	} else {
		fmt.Println("Not enough money.")
	}
}
