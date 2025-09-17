package player

import "fmt"

func IsDead(p *Player) {
	if p.Pv <= 0 {
		p.Pv = p.PvMax / 2
		fmt.Printf("\n GAME OVER \n")
		fmt.Printf("%s est mort mais ressuscite avec %d PV !\n", p.Name, p.Pv)
		fmt.Println("Vous perdez la moitié de vos roupies...")
		p.Money /= 2
	}
}
