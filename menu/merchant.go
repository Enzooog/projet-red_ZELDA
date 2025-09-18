package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

// menu display
func Merchant(p *player.Player) {
	fmt.Println("\n--- FARAS MERCHANT ---")
	fmt.Printf("Hello! Your rupees: %d\n", p.Money)
	fmt.Println("1. Sheikah armor (50 rupees) - Stealth")
	fmt.Println("2. Stone armor (50 rupees) - +20 HP")
	fmt.Println("3. Bird armor (50 rupees) - +10 speed")
	fmt.Println("4. Green potion (15 rupees) - Heals 30 HP")
	fmt.Println("5. Red potion  (20 rupees) - Enemy poison")
	fmt.Println("6. Bag (40 rupees) - +5 inventory slots")
	fmt.Println("7. Back")
	fmt.Print("Choice: ")

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

// Buy a bag to increase inventory capacity
func buyBag(p *player.Player) {
	if p.Money >= 40 {
		p.Money -= 40
		p.Inventory.MaxSize += 5
		fmt.Printf("Bag purchased! Inventory size: %d\n", p.Inventory.MaxSize)
	} else {
		fmt.Println("Not enough money.")
	}
}
