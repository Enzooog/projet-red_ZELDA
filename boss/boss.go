package boss

import (
	"fmt"
	"os"
	"time"

	"projet-red_ZELDA/combat"
	"projet-red_ZELDA/player"
)

// Structure representing a Boss
type Boss struct {
	Name   string // Boss name
	Pv     int    // Current health points
	PvMax  int    // Maximum health points
	Speed  int    // Speed (determines who attacks first)
	Damage int    // Damage dealt
}

// Creates the boss and starts the fight against the player
func CreateBoss(b *Boss, p *player.Player, e *combat.Enemy) {
	// Initialize the boss and enemy stats
	*b = Boss{Name: "Gannon", Pv: 200, PvMax: 200, Speed: 8, Damage: 20}
	*e = combat.Enemy{Name: b.Name, Pv: b.Pv, PvMax: b.PvMax, Speed: b.Speed, Damage: b.Damage}

	fmt.Printf(" %s appears!\n", b.Name)

	// Main fight loop
	for e.Pv > 0 && p.Pv > 0 {
		if p.Speed >= b.Speed {
			// Player attacks first if faster or equal speed
			combat.PlayerTurn(p, e)
			if e.Pv > 0 {
				combat.EnemyAttack(p, e) // Boss attack if alive
			}
		} else {
			// Boss attacks first if faster
			combat.EnemyAttack(p, e)
			if p.Pv > 0 {
				combat.PlayerTurn(p, e) // Player attack if alive
			}
		}
	}

	// End of the fight
	if e.Pv <= 0 {
		fmt.Println("WELL DONE, YOU BEAT THE BOSS!")
		time.Sleep(3 * time.Second) // Pause to read the message
		os.Exit(0)                  // End the game
	} else {
		player.IsDead(p) // Player is dead
	}
}
