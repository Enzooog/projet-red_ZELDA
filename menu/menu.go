package menu

import (
	"fmt"
	"os"
)

func menu() {
	var key int
	fmt.Scanf("%d", &key)

	switch key {
	case 1:
		key = 1
		fmt.Println("inventory")
	case 2:
		key = 2
		fmt.Println(" blacksmith ")
	case 3:
		key = 3
		fmt.Print("player information ")
	case 4:
		key = 4
		fmt.Print("resume the game")
	case 5:
		key = 5
		fmt.Print("exit the game ")
		var confirmation int
		fmt.Println("Are you sure you want to leave? ?")
		fmt.Println("1. yes")
		fmt.Println("2. No")
		fmt.Print("choice : ")
		fmt.Scan(&confirmation)

		if confirmation == 1 {
			fmt.Println("goodbye")
			os.Exit(0)
		} else if confirmation == 2 {
			fmt.Println("back to menu")
		} else {
			fmt.Println("invalid choice")
		}
	}
}
