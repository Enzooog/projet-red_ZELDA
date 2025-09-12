package menu

import "fmt"


func ConfirmExit() bool {
	var confirmation int
	fmt.Println("Are you sure you want to leave?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Print("Choice: ")
	fmt.Scan(&confirmation)

	return confirmation == 1
}
