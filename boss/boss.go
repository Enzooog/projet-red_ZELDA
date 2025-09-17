package boss

import (
	"fmt"

	"projet-red_ZELDA/player"
)

type Monster struct {
	Name        string
	AttackPower int
	PvMax       int
	Pv          int
}

func CreateBoss() Monster {
	return Monster{
		Name:        "ezz",
		AttackPower: 25,
		PvMax:       200,
		Pv:          200,
	}
}

func (m *Monster) Attack(p *player.Player) {
	fmt.Printf("%s attacks %s for %d damage!\n", m.Name, p.Name, m.AttackPower)
	p.Pv -= m.AttackPower
	if p.Pv < 0 {
		p.Pv = 0
	}
	fmt.Printf("%s now has %d/%d HP.\n", p.Name, p.Pv, p.PvMax)
}

func (m *Monster) TakeDamage(dmg int) {
	fmt.Printf("%s takes %d damage!\n", m.Name, dmg)
	m.Pv -= dmg
	if m.Pv < 0 {
		m.Pv = 0
	}
	fmt.Printf("%s now has %d/%d HP.\n", m.Name, m.Pv, m.PvMax)
}

func (m *Monster) IsDead() bool {
	if m.Pv <= 0 {
		fmt.Printf("%s has been defeated!\n", m.Name)
		return true
	}
	return false
}
