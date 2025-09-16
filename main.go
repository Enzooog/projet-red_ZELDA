package main

import (
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

	// Le joueur peut maintenant équiper un objet acheté
	p.Inventory.EquipItem()

	player.DisplayInfo(p)
}
