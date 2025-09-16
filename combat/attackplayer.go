package combat

import (
	"fmt"

	"projet-red_ZELDA/player"
)

func AttackPlayer(p *player.Player, damage int, e *Enemy) {
	p.Pv -= damage
	if p.Pv < 0 {
		p.Pv = 0
		AttackEnemy(e, 15, p)
	}
	fmt.Printf("%s t'attaque et t'inflige %d dégâts ! Tes PV restants : %d\n", e.Name, damage, p.Pv)
}
