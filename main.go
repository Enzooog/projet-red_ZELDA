package main

import (
	"fmt"


	"projet-red_ZELDA/combat"

	"projet-red_ZELDA/menu"
	"projet-red_ZELDA/player"
)

func main() {

	p := player.CreatePlayer() // Création du joueur directement

	for {
		fmt.Println("=== ZELDA GAME ===")
		fmt.Println("1. Menu")
		fmt.Println("2. Commencer les quêtes (battre l'enemy1)")

	p := player.CreatePlayer()
	player.DisplayInfo(p)

	p.Pv = 0
	player.IsDead(&p)
	player.TakePot(&p)
	for {
		fmt.Println("\nQue veux-tu faire ?")
		fmt.Println("1. Afficher infos joueur")
		fmt.Println("2. Acheter un objet ")
		fmt.Println("3. Afficher l'inventaire et équiper un objet")
		fmt.Println("4. Quitter")

		fmt.Print("Choix : ")
		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:

			MenuLoop(&p)
		case 2:
			StartQuest(&p)
		default:
			fmt.Println("Choix invalide.")
		}
	}
}


func MenuLoop(p *player.Player) {
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Inventaire")
		fmt.Println("2. Forgeron (à venir)")
		fmt.Println("3. Marchand")
		fmt.Println("4. Infos joueur")
		fmt.Println("5. Reprendre le jeu")
		fmt.Println("6. Quitter")
		fmt.Print("Choix : ")
		var key int
		fmt.Scan(&key)

		switch key {
		case 1:
			p.Inventory.ShowInventory()
		case 2:
			fmt.Println("Le forgeron n'est pas encore disponible.")
		case 3:
			menu.Merchant()
		case 4:
			player.DisplayInfo(*p)
		case 5:
			return // Reprendre le jeu
		case 6:
			if menu.ConfirmExit() {
				fmt.Println("Au revoir !")
				return
			}
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// Lancer la quête contre l'enemy1
func StartQuest(p *player.Player) {
	enemy := &combat.Enemy{Name: "Gobelin", Pv: 30, PvMax: 30}
	fmt.Printf("Un %s apparaît !\n", enemy.Name)
	for enemy.Pv > 0 && p.Pv > 0 {
		fmt.Println("Appuie sur Entrée pour attaquer...")
		fmt.Scanln()
		combat.AttackEnemy(enemy, 10, p)
		if enemy.Pv > 0 {
			combat.AttackPlayer(p, 5, enemy)
		}
	}
	if enemy.Pv <= 0 {
		fmt.Println("Tu as vaincu l'ennemi !")
	} else if p.Pv <= 0 {
		player.IsDead(p)
		fmt.Println("DEAD")

			player.DisplayInfo(p)
		case 2:
			menu.BuyStone(&p.Money, &p.Inventory)
		case 3:
			equipped := p.Inventory.EquipItem()
			if equipped != nil && equipped.BonusType == "pv" {
				p.PvMax += equipped.Bonus
				p.Pv += equipped.Bonus
				fmt.Printf("Bonus de PV appliqué : %d ! Nouveaux PV max : %d\n", equipped.Bonus, p.PvMax)
			}
		case 4:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}

		// Le joueur peut maintenant équiper un objet acheté
		p.Inventory.EquipItem()

		player.DisplayInfo(p)

	}
}
