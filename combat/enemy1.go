package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func PlayerTurn(p *player.Player, e *Enemy) {
	fmt.Println("\n--- Votre tour ---")
	fmt.Println("1. Attaque à l'épée")
	fmt.Println("2. Sort magique")
	fmt.Println("3. Utiliser une potion")
	fmt.Print("Choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		damage := p.SwordDmg
		e.Pv -= damage
		if e.Pv < 0 {
			e.Pv = 0
		}
		fmt.Printf("Vous attaquez avec l'épée ! %d dégâts. PV ennemi : %d/%d\n",
			damage, e.Pv, e.PvMax)
	case 2:
		damage := p.SwordDmg + 10
		e.Pv -= damage
		p.Pv -= 5
		if e.Pv < 0 {
			e.Pv = 0
		}
		if p.Pv < 0 {
			p.Pv = 0
		}
		fmt.Printf("Sort lancé ! %d dégâts à l'ennemi, 5 dégâts à vous-même.\n", damage)
		fmt.Printf("PV ennemi : %d/%d | Vos PV : %d/%d\n", e.Pv, e.PvMax, p.Pv, p.PvMax)
	case 3:
		player.UsePotion(p)
	}
}

func EnemyAttack(p *player.Player, e *Enemy) {
	p.Pv -= e.Damage
	if p.Pv < 0 {
		p.Pv = 0
	}
	fmt.Printf("%s vous attaque ! %d dégâts. Vos PV : %d/%d\n",
		e.Name, e.Damage, p.Pv, p.PvMax)
}
