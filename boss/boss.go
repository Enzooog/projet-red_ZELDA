package boss

import (
	"fmt"
	"projet-red_ZELDA/player"
	"projet-red_ZELDA/combat"
)

type Boss struct {
	Name   string
	Pv     int
	PvMax  int
	Speed  int
	Damage int
}

func CreateBoss(b *Boss, p *player.Player, e *combat.Enemy) {
	*b = Boss{Name: "Gannon", Pv: 200, PvMax: 200, Speed: 8, Damage: 20}
	*e = combat.Enemy{Name: b.Name, Pv: b.Pv, PvMax: b.PvMax, Speed: b.Speed, Damage: b.Damage}
	
	fmt.Printf("Un %s apparaît !\n", b.Name)

	for e.Pv > 0 && p.Pv > 0 {
		if p.Speed >= b.Speed {
			combat.PlayerTurn(p, e)
			if e.Pv > 0 {
				combat.EnemyAttack(p, e)
			}
		} else {
			combat.EnemyAttack(p, e)
			if p.Pv > 0 {
				combat.PlayerTurn(p, e)
			}
		}
	}

	if e.Pv <= 0 {
		fmt.Println("Victoire ! Vous gagnez 50 roupies !")
		p.Money += 50
	} else {
		player.IsDead(p)
	}
}