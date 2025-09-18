package player

import (
	"fmt"
	"time"
)

func ShowIntroLore() {
	lore := []string{
		"A hundred years ago, the kingdom of Hyrule was consumed by darkness.",
		"Calamity Ganon, an ancient evil, awakened and unleashed destruction upon the land.",
		"The Champions, the Guardians, and even the King himself fell in the battle to stop him.",
		"But hope was not lost...",
		"A hero, chosen by destiny, was placed in a deep slumber to recover.",
		"Now, the time has come. The sword that seals the darkness calls once more.",
		"You are must rise and reclaim your strength.",
		"Explore the wilds, face the shadows, and restore balance to Hyrule.",
	}

	for _, line := range lore {
		fmt.Println(line)
		time.Sleep(4 * time.Second) // Attente de 4 secondes entre chaque ligne
	}
}
