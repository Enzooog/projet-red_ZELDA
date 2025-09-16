package menu

import "fmt"

func buySheikah(Money *int) {
	fmt.Println("1. stealth mask for 10 roupies")
	fmt.Println("2. stealth suit for 20 roupies")
	fmt.Println("3. stealth boots for 20 roupies")

	var choice int
	fmt.Print("Choose an item: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		if *Money >= 10 {
			*Money -= 10
			fmt.Println("You bought the stealth mask!")
		} else {
			fmt.Println("Not enough money.")
		}
	case 2:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the stealth suit!")
		} else {
			fmt.Println("Not enough money.")
		}
	case 3:
		if *Money >= 20 {
			*Money -= 20
			fmt.Println("You bought the stealth boots!")
		} else {
			fmt.Println("Not enough money.")
		}
	default:
		fmt.Println("Invalid choice.")
	}

	fmt.Println("Remaining money:", *Money)
}
