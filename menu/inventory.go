package menu

import "projet-red_ZELDA/types"

// Function to display the player's inventory

func ShowPlayerInventory(inv *types.Inventory) {
	inv.ShowInventory() // call the func to show inventory
}
