package player

import "fmt"

func IsDead(p *Player) {
	if p.Pv <= 0 {
		p.Pv = p.PvMax / 2
		fmt.Printf("%s was dead but has been resurrected with %d HP!\n", p.Name, p.Pv)
	} else {
		fmt.Printf("%s is still alive with %d HP.\n", p.Name, p.Pv)
	}
}
