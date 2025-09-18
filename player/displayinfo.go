package player

import "fmt"

// Function to display player information in real time
func DisplayInfo(p Player) {
	fmt.Println("\n=== STATS DU JOUEUR ===")
	fmt.Printf("Name : %s\n", p.Name)
	fmt.Printf("Class : %s\n", p.Classe)
	fmt.Printf("level : %d\n", p.Level)
	fmt.Printf("PV : %d/%d\n", p.Pv, p.PvMax)
	fmt.Printf("speed : %d\n", p.Speed)
	fmt.Printf("Sword damage : %d\n", p.SwordDmg)
	fmt.Printf("Roupies : %d\n", p.Money)
	fmt.Printf("Inventory : %d/%d objects\n", len(p.Inventory.Items), p.Inventory.MaxSize)
}
