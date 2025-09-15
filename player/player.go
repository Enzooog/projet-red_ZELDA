package player

import (
	"fmt"
	"strings"
	"time"

	"projet-red_ZELDA/menu"
)

// Definition of the Player structure
type Player struct {
	Name      string
	Level     int
	Classe    string
	PvMax     int
	Money     int
	Inventory menu.Inventory
}

// Function to check if the class is valid
func isValidClass(input string) bool {
	validClasses := []string{"Warrior", "Mage", "Archer"}
	for _, class := range validClasses {
		if strings.EqualFold(input, class) {
			return true
		}
	}
	return false
}

// Function to create a player
func CreatePlayer() Player {
	var name string
	var classe string
	pv := 100

	fmt.Println("Welcome to Zelda")
	time.Sleep(1 * time.Second)
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	for {
		fmt.Print("Choose your class (Warrior, Mage, Archer): ")
		fmt.Scan(&classe)
		if isValidClass(classe) {
			break
		}
		fmt.Println("Invalid class. Please choose Warrior, Mage, or Archer.")
	}

	if classe == "Warrior" {
		pv += 20
	}
	if classe == "Mage" {
		pv += 0
	}
	if classe == "Archer" {
		pv -= 20
	}

	player := Player{
		Name:   name,
		Level:  1,
		Classe: classe,
		PvMax:  pv,
	}
	fmt.Printf("Hello %s, you are a %s.\n", player.Name, player.Classe)
	return player
}
