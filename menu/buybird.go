package menu

import "fmt"

func buyBird(Money *int,inv *Inventory) {
	fmt.Println("1. feather jewel for 10 roupies")
	fmt.Println("2. downy piaf tunic for 20 roupies")
	fmt.Println("3. fluffy piaf pants for 20 roupies")

	var choice int
	fmt.Print("Choose an item: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
			if *Money >= 10 {
			*Money -= 10
			fmt.Println("You bought the feather jewel")
			inv.AddItem(Item{Name: "feather jewel", BonusType: "speed", Bonus: 5})
		} else {
			fmt.Println("Not enough money.")
		}
	case 2:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the downy piaf tunic")
			inv.AddItem(Item{Name: "downy piaf tunic", BonusType: "speed", Bonus: 10})
		} else {
			fmt.Println("Not enough money.")
		}
	case 3:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the fluffy piaf pants")
			inv.AddItem(Item{Name: "fluffy piaf pants", BonusType: "speed", Bonus: 10})
		} else {
			fmt.Println("Not enough money.")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
