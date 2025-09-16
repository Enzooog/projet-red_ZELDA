package main

import (
	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {
	p := player.CreatePlayer()
	player.DisplayInfo(p)
	// Achat d’un objet avec l'argent du joueur
	menu.BuyStone(&p.Money, &p.Inventory)



	// Le joueur peut maintenant équiper un objet acheté
	p.Inventory.EquipItem()
}
