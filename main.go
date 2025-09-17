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
		fmt.Println("2. Quête principale (battre l'ennemi)")
		fmt.Println("3. Battre le boss")
		fmt.Print("Choix : ")

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
			fmt.Println("Choix invalide.")
		}
	}
}