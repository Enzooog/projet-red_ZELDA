package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
	"projet-red_ZELDA/types"
)

func buySheikah(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Armure Sheikah", Type: "armor", BonusType: "stealth", Bonus: 1}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
		}
	} else {
		fmt.Println("Pas assez d'argent.")
	}
}

func buyStone(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Armure Pierre", Type: "armor", BonusType: "pv", Bonus: 20}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.PvMax += 20
			p.Pv += 20
			fmt.Println("PV max augmentés !")
		}
	} else {
		fmt.Println("Pas assez d'argent.")
	}
}

func buyBird(p *player.Player) {
	if p.Money >= 50 {
		item := types.Item{Name: "Armure Oiseau", Type: "armor", BonusType: "speed", Bonus: 10}
		if p.Inventory.AddItem(item) {
			p.Money -= 50
			p.Speed += 10
			fmt.Println("Vitesse augmentée !")
		}
	} else {
		fmt.Println("Pas assez d'argent.")
	}
}
