package menu

import (
	"fmt"
	"os"
)

func ConfirmExit() bool {
	fmt.Println("\nÊtes-vous sûr de vouloir quitter ?")
	fmt.Println("1. Oui")
	fmt.Println("2. Non")
	fmt.Print("Choix : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Au revoir et merci d'avoir joué !")
		os.Exit(0)
		return true
	}
	return false
}
