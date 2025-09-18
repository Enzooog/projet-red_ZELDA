package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

// Structure representing an enemy
type Enemy struct {
	Name   string // Enemy name
	Pv     int    // Current health points
	PvMax  int    // Maximum health points
	Speed  int    // Speed (determines who attacks first)
	Damage int    // Damage dealt
}

// Starts a battle against a basic enemy
func StartBattle(p *player.Player) {
	// Create the enemy
	enemy := &Enemy{Name: "Goblin", Pv: 40, PvMax: 40, Speed: 8, Damage: 12}
	fmt.Printf("A %s appears!\n", enemy.Name)

	// Main battle loop
	for enemy.Pv > 0 && p.Pv > 0 {
		if p.Speed >= enemy.Speed {
			// Player attacks first if faster or equal
			PlayerTurn(p, enemy)
			if enemy.Pv > 0 {
				EnemyAttack(p, enemy) // Enemy retaliates if still alive
			}
		} else {
			// else enemy attacks first
			EnemyAttack(p, enemy)
			if p.Pv > 0 {
				PlayerTurn(p, enemy) // Player retaliates if still alive
			}
		}
	}

	// End of the battle
	if enemy.Pv <= 0 {
		fmt.Println("Victory! You win 30 rupees!")
		p.Money += 30 // Reward
	} else {
		player.IsDead(p) // Player is dead
	}
}
