package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

type Enemy struct {
	Name   string
	Pv     int
	PvMax  int
	Speed  int
	Damage int
}

func StartBattle(p *player.Player) {
	enemy := &Enemy{Name: "Gobelin", Pv: 40, PvMax: 40, Speed: 8, Damage: 12}
	fmt.Printf("one %s appears !\n", enemy.Name)

	for enemy.Pv > 0 && p.Pv > 0 {
		if p.Speed >= enemy.Speed {
			PlayerTurn(p, enemy)
			if enemy.Pv > 0 {
				EnemyAttack(p, enemy)
			}
		} else {
			EnemyAttack(p, enemy)
			if p.Pv > 0 {
				PlayerTurn(p, enemy)
			}
		}
	}

	if enemy.Pv <= 0 {
		fmt.Println("Victory! You win 30 roupies !")
		p.Money += 30
	} else {
		player.IsDead(p)
	}
}
