package main

import (
	"fmt"

	"projet-red_ZELDA/boss"
	"projet-red_ZELDA/combat"
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {
	p := player.CreatePlayer()
	b := &boss.Boss{}
	e := &combat.Enemy{}

	for {
		fmt.Println("\n=== ZELDA GAME ===")
		fmt.Println("1. Menu")
		fmt.Println("2. Main quest")
		fmt.Println("3. Defeat the boss")
		fmt.Print("Choice : ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			menu.ShowMenu(&p)
		case 2:
			combat.StartBattle(&p)
		case 3:
			boss.CreateBoss(b, &p, e)
		default:
			fmt.Println("invalid choice.")
		}
	}
}
