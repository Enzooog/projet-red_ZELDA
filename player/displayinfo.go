package player

import "fmt"

// Fonction pour afficher les infos du joueur
func DisplayInfo(p Player) {
	fmt.Println("\n Player Info")
	fmt.Println("Name:", p.Name)
	fmt.Println("Class:", p.Classe)
	fmt.Println("Level:", p.Level)
	fmt.Println("Max HP:", p.PvMax)
	fmt.Println("Current HP:", p.Pv, "/", p.PvMax)
	fmt.Println("Money:", p.Money)
}
