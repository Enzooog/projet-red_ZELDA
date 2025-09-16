package menu

import "fmt"

func Merchant() {
	var Money int
	fmt.Println("Hello, I'm Faras. Welcome to my shop. What would you like to buy? ")
	fmt.Println("your roupie = ", Money)
	fmt.Println("1. outfit sheikah for 50 roupies")
	fmt.Println("2. stone outfit for 50 roupies")
	fmt.Println("3. bird outfit for 50 roupies")
	fmt.Println("4. green potion (recover from life ) for 10 roupies")
	fmt.Println("5. red potion (throw poison)for 10 roupies")
	fmt.Println("6. bag for 20 roupies")

	var key2 int
	fmt.Print("Choice: ")
	fmt.Scan(&key2)

	switch key2 {
	case 1:
		buySheikah(&Money, &Inventory{})
	case 2:
		BuyStone(&Money, &Inventory{})
	case 3:
		buyBird(&Money, &Inventory{})
	case 4:
		// acheter la potion verte
	case 5:
		// acheter la potion rouge
	case 6:
		// acheter the bag
	}
}
