package Player

import (
	"fmt"
	"time"
)

type Playe struct {
	name     string
	age      int
	isRemote bool
}

func Player() {
	fmt.Print(" Welcome to Zelda")
	time.Sleep(4 * time.Second)
	fmt.Print(" Choose one breed from these three")
}
