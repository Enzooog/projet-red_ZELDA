package player

import "fmt"

func DisplayInfo(p Player) {
	fmt.Println("\n=== STATS DU JOUEUR ===")
	fmt.Printf("Nom : %s\n", p.Name)
	fmt.Printf("Classe : %s\n", p.Classe)
	fmt.Printf("Niveau : %d\n", p.Level)
	fmt.Printf("PV : %d/%d\n", p.Pv, p.PvMax)
	fmt.Printf("Vitesse : %d\n", p.Speed)
	fmt.Printf("Dégâts épée : %d\n", p.SwordDmg)
	fmt.Printf("Roupies : %d\n", p.Money)
	fmt.Printf("Inventaire : %d/%d objets\n", len(p.Inventory.Items), p.Inventory.MaxSize)
}
