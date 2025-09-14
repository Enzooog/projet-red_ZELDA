package Player

import (
	"fmt"
	"time"
)

type Playe struct {
	level int
}

func Player() {
	player1 := Playe{
		level: 1,
	}
	var name string
	var classe string
	fmt.Print(" Welcome to Zelda \n")
	time.Sleep(1 * time.Second)
	fmt.Print("Enter your name and choose one breed from these three: ")
	fmt.Scan(&name, &classe) // Reading input separated by space
	fmt.Printf("Hello %s, you are an %s.\n", name, classe)
	fmt.Println("Player level", player1.level)
}
