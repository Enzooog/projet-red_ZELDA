package Player

import (
	"fmt"
	"time"
)

type Playe struct {
	Name   string
	Level  int
	Classe string
	Pvmax  int
}

func Player() {
	fmt.Print(" Welcome to Zelda")
	time.Sleep(4 * time.Second)
	fmt.Print("Choose one breed from these three")
}

