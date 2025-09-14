package main

import (
	"projet-red_ZELDA/menu/player"
)

func main() {
	p := player.CreatePlayer()
	player.DisplayInfo(p)
}
