package menu

import "fmt"

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

	}
}
