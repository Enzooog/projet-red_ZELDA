package menu

import "fmt"

type Item struct {
	Name string
}

type Inventory struct {
	Items []Item
}

// Ajouter un item à l'inventaire
func (inv *Inventory) AddItem(item Item) {
	inv.Items = append(inv.Items, item)
	fmt.Printf("%s a été ajouté à l'inventaire.\n", item.Name)
}

// Afficher les items de l'inventaire
func (inv *Inventory) ShowInventory() {
	if len(inv.Items) == 0 {
		fmt.Println("L'inventaire est vide.")
		return
	}
	fmt.Println("Contenu de l'inventaire :")
	for i, item := range inv.Items {
		fmt.Printf("%d. %s\n", i+1, item.Name)
	}
}

// Permet au joueur d'équiper un objet de l'inventaire
func (inv *Inventory) EquipItem() {
	inv.ShowInventory()

	if len(inv.Items) == 0 {
		return
	}

	var choice int
	fmt.Print("Choisissez un objet à équiper (numéro) : ")
	fmt.Scan(&choice)

	if choice >= 1 && choice <= len(inv.Items) {
		fmt.Printf("Vous avez équipé : %s\n", inv.Items[choice-1].Name)
	} else {
		fmt.Println("Choix invalide.")
	}
}
