package main

import (
	"fmt"
	"projet-red_ZELDA/boss"
	"projet-red_ZELDA/combat"
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)
func clear() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	// Create the player
	p := player.CreatePlayer()
	// Initialize the boss and enemy
	b := &boss.Boss{}
	e := &combat.Enemy{}

	for {
		fmt.Println("\n=== ZELDA GAME ===")
		fmt.Println("1. Menu")
		fmt.Println("2. Main quest")
		fmt.Println("3. Defeat the boss")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			clear()
			menu.ShowMenu(&p) // Show the menu
			
		case 2:
			clear()
			combat.StartBattle(&p) // battle against an enemy
			
		case 3:
			clear()
			boss.CreateBoss(b, &p, e) // Fight the boss
		default:
			clear()
			fmt.Println("Invalid choice.")
		}
	}
}
