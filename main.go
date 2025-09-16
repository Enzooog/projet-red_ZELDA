package main

import (
	"fmt"
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {
	p := player.CreatePlayer()
	player.DisplayInfo(p)

	p.Pv = 0
	player.IsDead(&p)
	player.TakePot(&p)

	// Achat d’un objet avec l'argent du joueur
	menu.BuyStone(&p.Money, &p.Inventory)
	equipped := p.Inventory.EquipItem()
	if equipped != nil && equipped.BonusType == "pv" {
		p.PvMax += equipped.Bonus
		p.Pv += equipped.Bonus
		fmt.Printf("Bonus de PV appliqué :%d ! Nouveaux PV max : %d\n", p.Pv, p.PvMax)
	}
	// Le joueur peut maintenant équiper un objet acheté
	p.Inventory.EquipItem()

	player.DisplayInfo(p)
}
