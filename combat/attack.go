package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

// Player's turn in combat
func PlayerTurn(p *player.Player, e *Enemy) {
	fmt.Println("\n--- your turn ---")
	fmt.Println("1. sword stroke")
	fmt.Println("2. magic spell")
	fmt.Println("3. use a potion")
	fmt.Print("Choice: ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Sword attack
		damage := p.SwordDmg
		e.Pv -= damage
		if e.Pv < 0 {
			e.Pv = 0
		}
		fmt.Printf("You attack with the sword: %d damage. Enemy HP: %d/%d\n", damage, e.Pv, e.PvMax)

	case 2:
		// Magic spell attack (damages both enemy and player)
		damage := p.SwordDmg + 10
		e.Pv -= damage
		p.Pv -= 5
		if e.Pv < 0 {
			e.Pv = 0
		}
		if p.Pv < 0 {
			p.Pv = 0
		}
		fmt.Printf("Spell cast! %d damage to the enemy, 5 damage to yourself.\n", damage)
		fmt.Printf("Enemy HP: %d/%d | Your HP: %d/%d\n", e.Pv, e.PvMax, p.Pv, p.PvMax)

	case 3:
		// Use a healing potion
		player.UsePotion(p)
	}
}

// Enemy attacks the player
func EnemyAttack(p *player.Player, e *Enemy) {
	p.Pv -= e.Damage
	if p.Pv < 0 {
		p.Pv = 0
	}
	fmt.Printf("%s attacks you: %d damage. Your HP: %d/%d\n", e.Name, e.Damage, p.Pv, p.PvMax)
}
