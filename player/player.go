package player

import (
	"fmt"
	"time"
)

// Définition de la structure Player
type Player struct {
	Name   string
	Level  int
	Classe string
	PvMax  int
}

// Fonction pour créer un joueur
func CreatePlayer() Player {
	var name string
	var classe string

	fmt.Println("Welcome to Zelda")
	time.Sleep(1 * time.Second)
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Choose your class (Warrior, Mage, Archer): ")
	fmt.Scan(&classe)

	player := Player{
		Name:   name,
		Level:  1,
		Classe: classe,
		PvMax:  100,
	}

	fmt.Printf("Hello %s, you are a %s.\n", player.Name, player.Classe)
	fmt.Println("Player level:", player.Level)
	return player
}

// Fonction pour afficher les infos du joueur
func DisplayInfo(p Player) {
	fmt.Println("\n--- Player Info ---")
	fmt.Println("Name:", p.Name)
	fmt.Println("Class:", p.Classe)
	fmt.Println("Level:", p.Level)
	fmt.Println("Max HP:", p.PvMax)
}
