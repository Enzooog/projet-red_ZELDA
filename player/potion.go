package player

import (
	"fmt"
)

func UsePotion(p *Player) {
	potions := []int{}
	for i, item := range p.Inventory.Items {
		if item.Type == "potion" {
			potions = append(potions, i)
		}
	}

	if len(potions) == 0 {
		fmt.Println("Aucune potion disponible !")
		return
	}

	fmt.Println("Potions disponibles :")
	for i, potionIndex := range potions {
		item := p.Inventory.Items[potionIndex]
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}

	var choice int
	fmt.Print("Choisir une potion : ")
	fmt.Scan(&choice)

	if choice >= 1 && choice <= len(potions) {
		potionIndex := potions[choice-1]
		item := p.Inventory.Items[potionIndex]

		if item.BonusType == "heal" {
			heal := item.Bonus
			p.Pv += heal
			if p.Pv > p.PvMax {
				p.Pv = p.PvMax
			}
			fmt.Printf("Vous récupérez %d PV ! PV actuels : %d/%d\n", heal, p.Pv, p.PvMax)
		}

		p.Inventory.RemoveItem(potionIndex)
	}
}
