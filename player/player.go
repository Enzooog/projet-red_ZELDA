package player

import (
	"fmt"
	"strings"
	"time"

	"projet-red_ZELDA/types"
)

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
}

func CreatePlayer() Player {
	var name, classe string
	pv, speed, swordDmg := 100, 10, 15

	fmt.Println("Bienvenue dans Zelda")
	time.Sleep(1 * time.Second)
	fmt.Print("Entrez votre nom : ")
	fmt.Scan(&name)

	for {
		fmt.Print("Choisissez votre classe (Warrior, Mage, Archer) : ")
		fmt.Scan(&classe)
		if isValidClass(classe) {
			break
		}
		fmt.Println("Classe invalide.")
	}

	classeLower := strings.ToLower(classe)
	switch classeLower {
	case "warrior":
		pv += 20
	case "mage":
		pv += 0
	case "archer":
		pv -= 20
	}

	return Player{
		Name: name, Level: 1, Classe: classe, PvMax: pv, Pv: pv,
		Money: 100, Speed: speed, SwordDmg: swordDmg,
		Inventory: types.Inventory{Items: []types.Item{}, MaxSize: 5},
	}
}

func isValidClass(input string) bool {
	validClasses := []string{"Warrior", "Mage", "Archer"}
	for _, class := range validClasses {
		if strings.EqualFold(input, class) {
			return true
		}
	}
	return false
}
