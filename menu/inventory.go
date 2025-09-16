package menu

import (
	"projet-red_ZELDA/types"
)

// Alias pour utiliser les types depuis le package types
func ShowPlayerInventory(inv *types.Inventory) {
	inv.ShowInventory()
}
