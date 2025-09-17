package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func Blacksmith(p *player.Player) {
	fmt.Println("\n--- BLLACKSMITH ---")
	fmt.Printf("Hello %s ! I can improve your sword.\n", p.Name)
	fmt.Printf("Current damage : %d\n", p.SwordDmg)
	fmt.Printf("Your roupies : %d\n", p.Money)
	fmt.Println("1. Improvement +5 damage (30 roupies)")
	fmt.Println("2. Improvement +10 damage (60 roupies)")
	fmt.Println("3. Back")
	fmt.Print("choice : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if p.Money >= 30 {
			p.Money -= 30
			p.SwordDmg += 5
			fmt.Printf("Improved sword! New damage : %d\n", p.SwordDmg)
		} else {
			fmt.Println("Not enough money.")
		}
	case 2:
		if p.Money >= 60 {
			p.Money -= 60
			p.SwordDmg += 10
			fmt.Printf("Greatly improved sword! New damage : %d\n", p.SwordDmg)
		} else {
			fmt.Println("Not enough money.")
		}
	case 3:
		return
	default:
		fmt.Println("Invalid choice.")
	}
}
