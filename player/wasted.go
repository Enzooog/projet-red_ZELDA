package player

import "fmt"

func IsDead(p *Player) {
	if p.Pv <= 0 {
		p.Pv = p.PvMax / 2
		fmt.Printf("\n GAME OVER \n")
		fmt.Printf("%s has died but is resurrected with %d PV !\n", p.Name, p.Pv)
		fmt.Println("You lose half your rupees...")
		p.Money /= 2
	}
}
