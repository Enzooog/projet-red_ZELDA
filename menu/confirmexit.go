package menu

import (
	"fmt"
	"os"
)

func confirmExit() bool {
	var confirmation int
	fmt.Println("Are you sure you want to leave?")
	fmt.Println("1. Yes")
	fmt.Println("2. No")
	fmt.Print("Choice: ")
	fmt.Scan(&confirmation)
	fmt.Println("Goodbye ")
	os.Exit(0)
	return confirmation == 1
}
