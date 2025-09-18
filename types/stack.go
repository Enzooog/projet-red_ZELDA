package types

import "fmt"

// Displays the inventory contents
func (inv *Inventory) ShowInventory() {
	if len(inv.Items) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Printf("Inventory (%d/%d):\n", len(inv.Items), inv.MaxSize)
	for i, item := range inv.Items {
		equipped := ""
		quantity := ""
		if item.IsEquipped {
			equipped = " equipped"
		}
		if item.Quantity > 1 {
			quantity = fmt.Sprintf(" x%d", item.Quantity)
		}
		fmt.Printf("%d. %s%s%s\n", i+1, item.Name, quantity, equipped)
	}
}

// Removes an item from the inventory
func (inv *Inventory) RemoveItem(index int) {
	if index >= 0 && index < len(inv.Items) {
		inv.Items = append(inv.Items[:index], inv.Items[index+1:]...)
	}
}
