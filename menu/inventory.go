package menu

import (
	"fmt"
	"projet-red_ZELDA/types"
)

func ShowPlayerInventory(inv *types.Inventory) {
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	fmt.Printf("Inventaire (%d/%d) :\n", len(inv.Items), inv.MaxSize)
	for i, item := range inv.Items {
		equipped := ""
		if item.IsEquipped {
			equipped = " [ÉQUIPÉ]"
		}
		fmt.Printf("%d. %s%s\n", i+1, item.Name, equipped)
	}

	fmt.Println("1. Équiper")
	fmt.Println("2. Retour")
	fmt.Print("Choix : ")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		EquipItem(inv)
	case 2:
		return
	}
}

func EquipItem(inv *types.Inventory) {
	fmt.Print("Numéro de l'objet à équiper : ")
	var num int
	fmt.Scan(&num)

	if num >= 1 && num <= len(inv.Items) {
		newItem := &inv.Items[num-1]
		if newItem.Type == "armor" {
			// Déséquiper toutes les autres armures
			for i := range inv.Items {
				if inv.Items[i].Type == "armor" && inv.Items[i].IsEquipped {
					inv.Items[i].IsEquipped = false
				}
			}
			// Équiper la nouvelle
			newItem.IsEquipped = true
			fmt.Printf("%s équipé !\n", newItem.Name)
		}
	}
}
