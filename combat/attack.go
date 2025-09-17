package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func PlayerTurn(p *player.Player, e *Enemy) {
	fmt.Println("\n--- your turn ---")
	fmt.Println("1. sword stroke ")
	fmt.Println("2. magic spell ")
	fmt.Println("3. use a potion ")
	fmt.Print("Choice : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		damage := p.SwordDmg
		e.Pv -= damage
		if e.Pv < 0 {
			e.Pv = 0
		}
		fmt.Printf("You attack with the sword   %d damages. PV enemy : %d/%d\n",
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
		fmt.Printf("spell cast! %d damage to the enemy, 5 damage to yourself.\n", damage)
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
	fmt.Printf("%s attacks you. %d damage. Your PV : %d/%d\n", e.Name, e.Damage, p.Pv, p.PvMax)
}
