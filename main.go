package main

import (
	"fmt"
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {
	p := player.CreatePlayer()
	player.DisplayInfo(p)

	p.Pv -= 0
	fmt.Printf("Current HP: %d/%d\n", p.Pv, p.PvMax)

	// Achat d’un objet avec l'argent du joueur
	menu.BuyStone(&p.Money, &p.Inventory)

	// Le joueur peut maintenant équiper un objet acheté
	p.Inventory.EquipItem()
}
