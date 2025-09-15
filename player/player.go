package player

import (
	"fmt"
	"strings"
	"time"
)

// Definition of the Player structure
type Player struct {
	Name   string
	Level  int
	Classe string
	PvMax  int
}

// List of available classes
var validClasses = []string{"Warrior", "Mage", "Archer"}

// Function to check if the class is valid
func isValidClass(input string) bool {
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
