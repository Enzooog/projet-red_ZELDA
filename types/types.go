package types

import "fmt"

// Structure for item
type Item struct {
	Name       string
	Type       string
	BonusType  string
	Bonus      int
	IsEquipped bool
	Quantity   int
}

// Structure representing the inventory
type Inventory struct {
	Items   []Item // List of items
	MaxSize int    // Maximum inventory size
}

// Adds an item to the inventory
func (inv *Inventory) AddItem(item Item) bool {
	item.IsEquipped = false
	item.Quantity = 1

	// Stack potions if they already exist
	if item.Type == "potion" {
		for i := range inv.Items {
			if inv.Items[i].Name == item.Name {
				inv.Items[i].Quantity++
				fmt.Printf("%s x%d\n", item.Name, inv.Items[i].Quantity)
				return true
			}
		}
	}

	// Check if inventory is full
	if len(inv.Items) >= inv.MaxSize {
		fmt.Println("Full inventory!")
		return false
	}

	// Add the item
	inv.Items = append(inv.Items, item)
	fmt.Printf("%s added.\n", item.Name)
	return true
}
