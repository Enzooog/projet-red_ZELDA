package menu

import "fmt"

func BuyStone(Money *int, inv *Inventory) {
	fmt.Println("1. stone helmet for 10 roupies")
	fmt.Println("2. stone armor for 20 roupies")
	fmt.Println("3. stone boots for 20 roupies")

	var choice int
	fmt.Print("Choose an item: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if *Money >= 10 {
			*Money -= 10
			fmt.Println("You bought the stone helmet")
			inv.AddItem(Item{Name: "Stone Helmet", BonusType: "pv", Bonus: 5})
		} else {
			fmt.Println("Not enough money.")
		}
	case 2:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the stone armor")
			inv.AddItem(Item{Name: "Stone Armor", BonusType: "pv", Bonus: 10})
		} else {
			fmt.Println("Not enough money.")
		}
	case 3:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the stone boots")
			inv.AddItem(Item{Name: "Stone Boots", BonusType: "pv", Bonus: 10})
		} else {
			fmt.Println("Not enough money.")
		}
	default:
		fmt.Println("Invalid choice.")
	}
}
