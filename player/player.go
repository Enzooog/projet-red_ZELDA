package player

import (
	"fmt"
	"strings"
	"time"

	"projet-red_ZELDA/types"
)

// Structure with infos of the player
type Player struct {
	Name      string
	Level     int
	Classe    string
	PvMax     int
	Pv        int
	Money     int
	Speed     int
	SwordDmg  int
	Inventory types.Inventory
	Mana      int
	ManaMax   int
}

// Creates player
func CreatePlayer() Player {
	var name, classe string
	pv, speed, swordDmg := 100, 10, 15

	fmt.Println("WELCOME TO ZELDA")
	time.Sleep(1 * time.Second)
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	for {
		fmt.Print("Choose your class (Warrior, Mage, Archer): ")
		fmt.Scan(&classe)
		if isValidClass(classe) {
			break
		}
		fmt.Println("Invalid class.")
	}

	classeLower := strings.ToLower(classe)
	switch classeLower {
	case "warrior":
		pv += 20 // HP bonus for Warrior
	case "mage":
		pv += 0 // No bonus for Mage
	case "archer":
		pv -= 20 // HP penalty for Archer
	}

	return Player{
		Name:      name,
		Level:     1,
		Classe:    classe,
		PvMax:     pv,
		Pv:        pv,
		Money:     100,
		Speed:     speed,
		SwordDmg:  swordDmg,
		Inventory: types.Inventory{Items: []types.Item{}, MaxSize: 5},
		Mana:      20,
		ManaMax:   40,
	}
}

// Checks if the chosen class is valid
func isValidClass(input string) bool {
	validClasses := []string{"Warrior", "Mage", "Archer"}
	for _, class := range validClasses {
		if strings.EqualFold(input, class) {
			return true
		}
	}
	return false
}
