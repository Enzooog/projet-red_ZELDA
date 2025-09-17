package types

import "fmt"

type Item struct {
	Name      string
	Type      string
	BonusType string
	Bonus     int
	IsEquipped bool
	Quantity  int
}

type Inventory struct {
	Items   []Item
	MaxSize int
}

func (inv *Inventory) AddItem(item Item) bool {
	item.IsEquipped = false
	item.Quantity = 1
	
	// Stack les potions
	if item.Type == "potion" {
		for i := range inv.Items {
			if inv.Items[i].Name == item.Name {
				inv.Items[i].Quantity++
				fmt.Printf("%s x%d\n", item.Name, inv.Items[i].Quantity)
				return true
			}
		}
	}
	
	if len(inv.Items) >= inv.MaxSize {
		fmt.Println("Full inventory !")
		return false
	}
	inv.Items = append(inv.Items, item)
	fmt.Printf("%s added.\n", item.Name)
	return true
}