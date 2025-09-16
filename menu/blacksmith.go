package menu

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func Blacksmith(p *player.Player) {
	fmt.Println("\n--- FORGERON ---")
	fmt.Printf("Bonjour %s ! Je peux améliorer votre épée.\n", p.Name)
	fmt.Printf("Dégâts actuels : %d\n", p.SwordDmg)
	fmt.Printf("Vos roupies : %d\n", p.Money)
	fmt.Println("1. Amélioration +5 dégâts (30 roupies)")
	fmt.Println("2. Amélioration +10 dégâts (60 roupies)")
	fmt.Println("3. Retour")
	fmt.Print("Choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if p.Money >= 30 {
			p.Money -= 30
			p.SwordDmg += 5
			fmt.Printf("Épée améliorée ! Nouveaux dégâts : %d\n", p.SwordDmg)
		} else {
			fmt.Println("Pas assez d'argent.")
		}
	case 2:
		if p.Money >= 60 {
			p.Money -= 60
			p.SwordDmg += 10
			fmt.Printf("Épée grandement améliorée ! Nouveaux dégâts : %d\n", p.SwordDmg)
		} else {
			fmt.Println("Pas assez d'argent.")
		}
	case 3:
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
