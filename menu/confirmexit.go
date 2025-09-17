package menu

import (
	"fmt"
	"os"
)

func ConfirmExit() bool {
	fmt.Println("\nAre you sure you want to leave ?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Print("Choice : ")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Goodbye and thank you for playing. !")
		os.Exit(0)
		return true
	}
	return false
}
