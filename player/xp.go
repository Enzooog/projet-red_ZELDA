package player

import "fmt"

func GainXP(p *Player) {
	for p.XP >= p.Level*100 {
		p.XP -= p.Level * 100
		p.Level++
		p.PvMax += 5
		p.ManaMax += 5
		p.SwordDmg += 2
		p.Pv = p.PvMax
		p.Mana = p.ManaMax
		fmt.Printf("Level up! You are now level %d!\n", p.Level)
		fmt.Printf("Stats increased: HP %d, Mana %d, Sword Damage %d\n", p.PvMax, p.ManaMax, p.SwordDmg)
	}
}
