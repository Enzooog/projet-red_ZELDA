package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

type Enemy struct {
	Name  string
	Pv    int
	PvMax int
}

func AttackEnemy(e *Enemy, damage int, p *player.Player) {
	e.Pv -= damage
	if e.Pv < 0 {
		e.Pv = 0
		AttackPlayer(p, 15, e)
	}
	fmt.Printf("Tu attaques %s et lui infliges %d dégâts ! PV restants : %d\n", e.Name, damage, e.Pv)
}
