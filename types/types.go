package types

import "fmt"

type Item struct {
	Name      string
	Type      string 
	BonusType string 
	Bonus     int
}

type Inventory struct {
	Items   []Item
	MaxSize int
}

func (inv *Inventory) AddItem(item Item) bool {
	if len(inv.Items) >= inv.MaxSize {
		fmt.Println("Inventaire plein !")
		return false
	}
	inv.Items = append(inv.Items, item)
	fmt.Printf("%s ajouté à l'inventaire.\n", item.Name)
	return true
}

func (inv *Inventory) ShowInventory() {
	if len(inv.Items) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Printf("Inventaire (%d/%d) :\n", len(inv.Items), inv.MaxSize)
	for i, item := range inv.Items {
		bonus := ""
		if item.Bonus > 0 {
			bonus = fmt.Sprintf(" (+%d %s)", item.Bonus, item.BonusType)
		}
		fmt.Printf("%d. %s%s\n", i+1, item.Name, bonus)
	}
}

func (inv *Inventory) RemoveItem(index int) {
	if index >= 0 && index < len(inv.Items) {
		inv.Items = append(inv.Items[:index], inv.Items[index+1:]...)
	}
}
