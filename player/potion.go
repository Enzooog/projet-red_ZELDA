package player

import "fmt"

// Use a potion to recover HP
func UsePotion(p *Player) {
	potions := []int{}
	// Find all potions in the inventory
	for i, item := range p.Inventory.Items {
		if item.Type == "potion" {
			potions = append(potions, i)
		}
	}

	// Check if there are available potions
	if len(potions) == 0 {
		fmt.Println("No potions available!")
		return
	}

	fmt.Println("Available potions:")
	for i, potionIndex := range potions {
		item := p.Inventory.Items[potionIndex]
		fmt.Printf("%d. %s x%d\n", i+1, item.Name, item.Quantity)
	}

	var choice int
	fmt.Scan(&choice)

	// Check player choice
	if choice >= 1 && choice <= len(potions) {
		potionIndex := potions[choice-1]
		item := &p.Inventory.Items[potionIndex]

		// Recover HP
		p.Pv += item.Bonus
		if p.Pv > p.PvMax {
			p.Pv = p.PvMax
		}
		fmt.Printf("You recover %d HP!\n", item.Bonus)

		// Decrease potion quantity and remove if empty
		item.Quantity--
		if item.Quantity <= 0 {
			p.Inventory.RemoveItem(potionIndex)
		}
	}
}
