package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/term"
)

func exploreMap(player *Character, reader *bufio.Reader) bool {
	cooldownDuration := 120 * time.Second
	timeSinceLastExplore := time.Since(player.LastExploreTime)

	if timeSinceLastExplore < cooldownDuration {
		remaining := cooldownDuration - timeSinceLastExplore
		clearScreen()
		fmt.Println(Cyan + "--- 🗺️ EXPLORER YNOV ---" + Reset)
		fmt.Println(Yellow + fmt.Sprintf("\n⏳ Cooldown actif ! Tu dois encore attendre %.0f secondes avant de pouvoir explorer Ynov à nouveau.", remaining.Seconds()) + Reset)
		fmt.Print("\n[Appuyez sur Entrée pour revenir]")
		reader.ReadString('\n')
		return false
	}

	player.LastExploreTime = time.Now()

	gridSize := 15
	px, py := 0, 0
	ex, ey := 14, 14

	events := make(map[string]string)
	visited := make(map[string]bool)
	visited["0,0"] = true

	for i := 0; i < 20; i++ {
		rx, ry := rand.Intn(gridSize), rand.Intn(gridSize)
		if (rx != 0 || ry != 0) && (rx != ex || ry != ey) {
			events[fmt.Sprintf("%d,%d", rx, ry)] = "monster"
		}
	}
	for i := 0; i < 12; i++ {
		rx, ry := rand.Intn(gridSize), rand.Intn(gridSize)
		if (rx != 0 || ry != 0) && (rx != ex || ry != ey) {
			events[fmt.Sprintf("%d,%d", rx, ry)] = "chest"
		}
	}
	for i := 0; i < 8; i++ {
		rx, ry := rand.Intn(gridSize), rand.Intn(gridSize)
		if (rx != 0 || ry != 0) && (rx != ex || ry != ey) {
			events[fmt.Sprintf("%d,%d", rx, ry)] = "trap"
		}
	}

	StatusMsg = "Vous entrez dans le campus Ynov..."

	for {
		clearScreen()
		fmt.Println(Cyan + "╭──────────────────────────────────────────────────────────────────╮" + Reset)
		fmt.Println(Cyan + "│          🗺️ EXPLORER YNOV - CAMPUS TACTIQUE (15x15)              │" + Reset)
		fmt.Println(Cyan + "╰──────────────────────────────────────────────────────────────────╯" + Reset)

		fmt.Println("    ┌" + strings.Repeat("───", gridSize) + "┐")
		for y := 0; y < gridSize; y++ {
			fmt.Print("    │")
			for x := 0; x < gridSize; x++ {
				posKey := fmt.Sprintf("%d,%d", x, y)

				if x == px && y == py {
					fmt.Print(Green + " 🧍" + Reset)
				} else if visited[posKey] {
					fmt.Print(DarkGray + "  ·" + Reset)
				} else {
					fmt.Print(Gray + "  ░" + Reset)
				}
			}
			fmt.Println("│")
		}
		fmt.Println("    └" + strings.Repeat("───", gridSize) + "┘")

		fmt.Println(Gray + "──────────────────────────────────────────────────────────────────" + Reset)
		fmt.Printf(" ❤️ PV : %d/%d  |  ⭐ Magie : %d/%d  |  🏆 Victoires : %d/15\n", player.CurrentHP, player.MaxHP, player.CurrentMagie, player.MaxMagie, player.FightsWon)
		fmt.Printf(" 📡 STATUT : %s\n", StatusMsg)
		fmt.Println(Gray + "──────────────────────────────────────────────────────────────────" + Reset)

		fmt.Println(" 🎮 Commandes : [Z] Haut | [S] Bas | [Q] Gauche | [D] Droite | [0] Quitter")
		fmt.Print("► Direction : ")

		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			break
		}

		buf := make([]byte, 1)
		os.Stdin.Read(buf)
		term.Restore(int(os.Stdin.Fd()), oldState)

		dir := strings.ToUpper(string(buf[0]))

		if dir == "0" || dir == "\x03" {
			break
		}

		nx, ny := px, py
		switch dir {
		case "Z":
			ny--
		case "S":
			ny++
		case "Q":
			nx--
		case "D":
			nx++
		}

		if nx >= 0 && nx < gridSize && ny >= 0 && ny < gridSize {
			px, py = nx, ny
			posKey := fmt.Sprintf("%d,%d", px, py)
			visited[posKey] = true

			if px == ex && py == ey {
				StatusMsg = "🎉 Vous avez atteint la sortie d'Ynov !"
				clearScreen()
				fmt.Println(Green + "\n" + StatusMsg + Reset)
				time.Sleep(1500 * time.Millisecond)
				break
			}

			ev := events[posKey]
			if ev == "monster" {
				StatusMsg = "⚠️ Une menace surgit des couloirs ! Combat !"
				clearScreen()
				fmt.Println(Red + "\n" + StatusMsg + Reset)
				time.Sleep(700 * time.Millisecond)
				enemy := generateRandomMonster(player.Level, player.Exp)
				won := runFight(player, &enemy, reader)
				if player.FightsWon >= 15 {
					return true
				}
				if player.CurrentHP <= 0 {
					return false
				}
				if won {
					StatusMsg = "Vous reprenez votre marche après l'affrontement victorieux."
				} else {
					StatusMsg = "Fuite effectuée... Vous êtes une peureuse !"
				}
			} else if ev == "chest" {
				coins := rand.Intn(20) + 10
				player.Credits += coins
				StatusMsg = fmt.Sprintf("💰 Vous découvrez un casier contenant +%d 🪙 !", coins)
			} else if ev == "trap" {
				dmg := rand.Intn(10) + 5
				player.CurrentHP -= dmg
				if player.CurrentHP < 0 {
					player.CurrentHP = 0
				}
				StatusMsg = fmt.Sprintf("💥 Câble réseau défectueux ! -%d ❤️.", dmg)
			} else {
				StatusMsg = "Vous avancez prudemment. Le campus semble calme..."
			}
		} else {
			StatusMsg = "❌ Déplacement impossible : Fin du campus atteinte."
		}
		time.Sleep(100 * time.Millisecond)
	}
	return player.FightsWon >= 15
}
