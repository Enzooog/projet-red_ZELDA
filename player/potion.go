package player

import "fmt"

func UsePotion(p *Player) {
	potions := []int{}
	for i, item := range p.Inventory.Items {
		if item.Type == "potion" {
			potions = append(potions, i)
		}
	}

	if len(potions) == 0 {
		fmt.Println("No potions available !")
		return
	}

	fmt.Println("Available potions :")
	for i, potionIndex := range potions {
		item := p.Inventory.Items[potionIndex]
		fmt.Printf("%d. %s x%d\n", i+1, item.Name, item.Quantity)
	}

	var choice int
	fmt.Scan(&choice)

	if choice >= 1 && choice <= len(potions) {
		potionIndex := potions[choice-1]
		item := &p.Inventory.Items[potionIndex]

		p.Pv += item.Bonus
		if p.Pv > p.PvMax {
			p.Pv = p.PvMax
		}
		fmt.Printf("You recover %d PV !\n", item.Bonus)

		item.Quantity--
		if item.Quantity <= 0 {
			p.Inventory.RemoveItem(potionIndex)
		}
	}
}
