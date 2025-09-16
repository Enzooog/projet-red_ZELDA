package player

import "fmt"

func TakePot(p *Player) {
	if p.Pv == p.PvMax {
		fmt.Println("You are already at full health. Potion not needed.")
		return
	}

	heal := 30
	p.Pv += heal
	if p.Pv > p.PvMax {
		p.Pv = p.PvMax
	}
	fmt.Printf("You used a potion and recovered %d HP. Current HP : %d/%d\n", heal, p.Pv, p.PvMax)
}
