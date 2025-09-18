package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

// Function for the Blacksmith menu
func Blacksmith(p *player.Player) {
	fmt.Println("\n--- BLACKSMITH ---")
	fmt.Printf("Hello %s! I can improve your sword.\n", p.Name)
	fmt.Printf("Current damage: %d\n", p.SwordDmg)
	fmt.Printf("Your rupees: %d\n", p.Money)
	fmt.Println("1. Improve +5 damage (30 rupees)")
	fmt.Println("2. Improve +10 damage (60 rupees)")
	fmt.Println("3. Back")
	fmt.Print("Choice: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Small sword improvement
		if p.Money >= 30 {
			p.Money -= 30
			p.SwordDmg += 5
			fmt.Printf("Sword improved! New damage: %d\n", p.SwordDmg)
		} else {
			fmt.Println("Not enough money.")
		}
	case 2:
		// Bigger sword improvement
		if p.Money >= 60 {
			p.Money -= 60
			p.SwordDmg += 10
			fmt.Printf("Sword greatly improved! New damage: %d\n", p.SwordDmg)
		} else {
			fmt.Println("Not enough money.")
		}
	case 3:
		// Go back to the previous menu
		return
	default:
		fmt.Println("Invalid choice.")
	}
}
