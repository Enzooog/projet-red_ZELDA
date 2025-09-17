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
	boss := &Boss{Name: "Gannon", Pv: 40, PvMax: 40, Speed: 8, Damage: 12}
	fmt.Printf("Un %s apparaît !\n", boss.Name)

	for boss.Pv > 0 && p.Pv > 0 {
		if p.Speed >= boss.Speed {
			combat.PlayerTurn(p,e)
			if boss.Pv > 0 {
				CreateBoss(b, p, e) 
			}
		} else {
			CreateBoss(b, p, e) 
			if p.Pv > 0 {
				combat.PlayerTurn(p, e) 
			}
		}
	}

	if boss.Pv <= 0 {
		fmt.Println("Victoire ! Vous gagnez 30 roupies !")
		p.Money += 50 
	} else {
		player.IsDead(p)
	}
}