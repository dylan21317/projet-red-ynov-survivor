package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Fonction pour charger et retourner l'art ASCII du boss depuis ascii.txt
func getMentorSprite() []string {
	file, err := os.Open("ascii.txt")
	if err != nil {
		// Sprite de secours si le fichier ascii.txt n'est pas trouvé
		return []string{
			"    [MENTOR SUPRÊME]    ",
			"       ( •̀_•́ )          ",
			"      /|#####|\\         ",
			"     / |     | \\        ",
			"    /  |     |  \\       ",
		}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getHeroSprite(class string) []string {
	switch class {
	case "Elfe":
		return []string{
			"      /\\      ",
			"     /  \\     ",
			"    ( o.o )   ",
			"   /|  |\\     ",
			"  / |  | \\    ",
			"    /  \\      ",
			"   /    \\     ",
		}
	case "Nain":
		return []string{
			"   [=======]  ",
			"   ( •̀-•́ )  ",
			"   <| ### |>  ",
			"    /     \\   ",
			"   | || || |  ",
			"    /     \\   ",
			"   /       \\  ",
		}
	default:
		return []string{
			"      ___     ",
			"     / o \\    ",
			"     \\___/    ",
			"     / | \\    ",
			"    /  |  \\   ",
			"       |      ",
			"      / \\     ",
		}
	}
}

func getEnemySprite(enemyType string) []string {
	switch enemyType {
	case "drone":
		return []string{
			"    _ [o=o] _ ",
			"  /  /===\\  \\",
			" |  |=====|  |",
			"  \\  \\___/  /",
			"   *  *  *    ",
			"  / /   \\ \\   ",
			"             ",
		}
	case "cyborg":
		return []string{
			"    _______   ",
			"   / (X_X) \\  ",
			"  <|  X_X  |> ",
			"   \\  |v|  /  ",
			"   / /   \\ \\  ",
			"  / /     \\ \\ ",
			"             ",
		}
	case "mentor":
		return getMentorSprite() // <--- Appelle votre fichier ascii.txt ici !
	default:
		return []string{
			"     .-.      ",
			"    (o.O)     ",
			"   /|; ;|\\    ",
			"  / |===| \\   ",
			"    /   \\     ",
			"   /     \\    ",
			"             ",
		}
	}
}

func generateRandomMonster(playerLevel int, playerExp int) Enemy {
	if FledEnemy != nil {
		enemyToFight := *FledEnemy
		FledEnemy = nil
		return enemyToFight
	}

	monsters := []Enemy{
		{Name: "Drone Sentinelle", MaxHP: 50, CurrentHP: 50, Attack: 8, Initiative: 12, ExpReward: 55, Type: "drone"},
		{Name: "Cyborg Rejeté", MaxHP: 70, CurrentHP: 70, Attack: 11, Initiative: 7, ExpReward: 70, Type: "cyborg"},
		{Name: "Mutant Acide", MaxHP: 60, CurrentHP: 60, Attack: 9, Initiative: 9, ExpReward: 60, Type: "mutant"},
	}

	baseMonster := monsters[rand.Intn(len(monsters))]
	scalingFactor := float64(playerLevel-1)*0.3 + float64(playerExp)/300.0

	baseMonster.MaxHP = int(float64(baseMonster.MaxHP) * (1.0 + scalingFactor))
	baseMonster.CurrentHP = baseMonster.MaxHP
	baseMonster.Attack = int(float64(baseMonster.Attack) * (1.0 + scalingFactor*0.6))
	baseMonster.Initiative += int(scalingFactor * 2)

	return baseMonster
}

func drawArenaFrame(player *Character, enemy *Enemy, combatLog string, animStep int, projIcon string, isPlayerAttacking bool) {
	clearScreen()
	fmt.Println(Cyan + "--- ⚔️ ZONE DE COMBAT TACTIQUE (YNOV) ---" + Reset)
	burnStatusP := ""
	if player.IsBurning {
		burnStatusP = Red + " 🔥(BRÛLURE)" + Reset
	}
	burnStatusE := ""
	if enemy.IsBurning {
		burnStatusE = Red + " 🔥(BRÛLURE)" + Reset
	}

	fmt.Printf(" Victoires : %d 🏆 | %s%s%s%s (❤️ %d/%d | ⭐ %d/%d)  VS  %s%s%s%s (❤️ %d/%d)\n", player.FightsWon, Green+Bold, player.Name, Reset, burnStatusP, player.CurrentHP, player.MaxHP, player.CurrentMagie, player.MaxMagie, Red+Bold, enemy.Name, Reset, burnStatusE, enemy.CurrentHP, enemy.MaxHP)
	fmt.Println(Gray + "------------------------------------------------------------------" + Reset)

	heroSprite := getHeroSprite(player.Class)
	enemySprite := getEnemySprite(enemy.Type)
	trackWidth := 14

	// On prend la hauteur maximale entre le héros et le boss pour ne rien couper
	maxLines := len(heroSprite)
	if len(enemySprite) > maxLines {
		maxLines = len(enemySprite)
	}

	for i := 0; i < maxLines; i++ {
		line := "    "

		if i < len(heroSprite) {
			line += fmt.Sprintf("%s%s%s", Green, heroSprite[i], Reset)
		} else {
			line += strings.Repeat(" ", 15)
		}

		if i == 2 && projIcon != "" {
			leftPad := 0
			rightPad := 0

			if isPlayerAttacking {
				leftPad = animStep
				rightPad = trackWidth - animStep - 2
			} else {
				leftPad = trackWidth - animStep - 2
				rightPad = animStep
			}

			if leftPad < 0 {
				leftPad = 0
			}
			if rightPad < 0 {
				rightPad = 0
			}

			line += strings.Repeat(" ", leftPad) + projIcon + strings.Repeat(" ", rightPad)
		} else {
			line += strings.Repeat(" ", trackWidth)
		}

		if i < len(enemySprite) {
			line += fmt.Sprintf("%s%s%s", Red, enemySprite[i], Reset)
		}
		fmt.Println(line)
	}

	fmt.Println(Gray + "------------------------------------------------------------------" + Reset)
	if combatLog != "" {
		fmt.Printf(" 💬 %s%s%s\n", Yellow, combatLog, Reset)
	} else {
		fmt.Println(" 💬 Analyse de la situation tactique...")
	}
	fmt.Println(Gray + "------------------------------------------------------------------" + Reset)
}

func animateAttack(player *Character, enemy *Enemy, combatLog string, icon string, isPlayerAttacking bool) {
	steps := 12
	for s := 0; s <= steps; s += 3 {
		drawArenaFrame(player, enemy, combatLog, s, icon, isPlayerAttacking)
		time.Sleep(40 * time.Millisecond)
	}
	drawArenaFrame(player, enemy, combatLog, steps, "💥", isPlayerAttacking)
	time.Sleep(100 * time.Millisecond)
}

func checkBurnsInCombat(player *Character, enemy *Enemy) string {
	msg := ""
	now := time.Now()

	if enemy.IsBurning {
		if now.Before(enemy.BurnTimer) {
			dmg := 4
			enemy.CurrentHP -= dmg
			if enemy.CurrentHP < 0 {
				enemy.CurrentHP = 0
			}
			msg += fmt.Sprintf(" %s souffre de la brûlure (-%d ❤️).", enemy.Name, dmg)
		} else {
			enemy.IsBurning = false
		}
	}

	if player.IsBurning {
		if now.Before(player.BurnTimer) {
			dmg := 4
			player.CurrentHP -= dmg
			if player.CurrentHP < 0 {
				player.CurrentHP = 0
			}
			msg += fmt.Sprintf(" Vous brûlez sous l'effet de la chaleur (-%d ❤️).", dmg)
		} else {
			player.IsBurning = false
		}
	}
	return msg
}

func runFight(player *Character, enemy *Enemy, reader *bufio.Reader) bool {
	combatLog := "Alerte : Engagement direct avec " + enemy.Name + " !"
	enemyStunned := false

	for player.CurrentHP > 0 && enemy.CurrentHP > 0 {
		burnMsg := checkBurnsInCombat(player, enemy)
		if burnMsg != "" {
			combatLog += burnMsg
		}

		if player.CurrentHP <= 0 || enemy.CurrentHP <= 0 {
			break
		}

		drawArenaFrame(player, enemy, combatLog, 0, "", true)

		fmt.Println("\n Actions : [1] Compétences & Sorts  [2] Inventaire  [0] Fuir (-10 ⭐)")
		fmt.Print("► Choix : ")
		actIn, _ := reader.ReadString('\n')
		act := strings.TrimSpace(actIn)

		if act == "2" {
			clearScreen()
			if len(player.Inventory) == 0 {
				combatLog = "🎒 Sac à dos vide !"
				continue
			}
			fmt.Println("\n-- Inventaire en combat --")
			for i, it := range player.Inventory {
				fmt.Printf(" [%d] %s\n", i+1, it)
			}
			fmt.Print("► Numéro de l'objet : ")
			itemIn, _ := reader.ReadString('\n')
			val, _ := strconv.Atoi(strings.TrimSpace(itemIn))

			if val >= 1 && val <= len(player.Inventory) {
				it := player.Inventory[val-1]
				if it == "🍷 Potion de soin basique" {
					player.CurrentHP += 20
					if player.CurrentHP > player.MaxHP {
						player.CurrentHP = player.MaxHP
					}
					removeItem(player, it)
					combatLog = "✨ Potion de soin basique injectée (+20 ❤️)."
				} else if it == "🍷 Potion de soin supérieure" {
					player.CurrentHP += 50
					if player.CurrentHP > player.MaxHP {
						player.CurrentHP = player.MaxHP
					}
					removeItem(player, it)
					combatLog = "✨ Potion de soin supérieure injectée (+50 ❤️)."
				} else if it == "🧪 Potion de magie basique" {
					player.CurrentMagie += 20
					if player.CurrentMagie > player.MaxMagie {
						player.CurrentMagie = player.MaxMagie
					}
					removeItem(player, it)
					combatLog = "✨ Potion de magie basique injectée (+20 ⭐)."
				} else if it == "🧪 Potion de magie supérieure" {
					player.CurrentMagie += 40
					if player.CurrentMagie > player.MaxMagie {
						player.CurrentMagie = player.MaxMagie
					}
					removeItem(player, it)
					combatLog = "✨ Potion de magie supérieure injectée (+40 ⭐)."
				} else {
					combatLog = "Cet objet ne peut pas être utilisé en combat."
					continue
				}
			} else {
				continue
			}
		} else if act == "1" {
			clearScreen()
			fmt.Println("\n-- Arsenal disponible --")
			for i, sk := range player.Skills {
				cdText := ""
				if sk.CurrentCD > 0 {
					cdText = fmt.Sprintf(" [Recharge: %d tour(s)]", sk.CurrentCD)
				}
				fmt.Printf(" [%d] %s %-20s (Dégâts: %2d | Coût: %2d ⭐)%s\n", i+1, sk.Icon, sk.Name, sk.Damage, sk.MagieCost, cdText)
			}
			fmt.Print("► Choix de l'attaque : ")
			skIn, _ := reader.ReadString('\n')
			val, _ := strconv.Atoi(strings.TrimSpace(skIn))

			if val >= 1 && val <= len(player.Skills) {
				sk := &player.Skills[val-1]

				if sk.CurrentCD > 0 {
					combatLog = fmt.Sprintf("⏳ Ce sort est en recharge (%d tour(s) restant(s)) !", sk.CurrentCD)
					continue
				}

				if player.CurrentMagie < sk.MagieCost {
					combatLog = "❌ Énergie magique (⭐) insuffisante !"
					continue
				}

				player.CurrentMagie -= sk.MagieCost
				sk.CurrentCD = sk.CooldownMax
				combatLog = fmt.Sprintf("Vous déclenchez : %s !", sk.Name)

				animateAttack(player, enemy, combatLog, sk.Icon, true)

				enemy.CurrentHP -= sk.Damage
				if enemy.CurrentHP < 0 {
					enemy.CurrentHP = 0
				}
				combatLog = fmt.Sprintf("Impact ! %s inflige %d dégâts.", sk.Name, sk.Damage)

				if sk.EffectType == "brulure" {
					enemy.IsBurning = true
					enemy.BurnTimer = time.Now().Add(2 * time.Second)
					combatLog += " 🔥 Enflammé ! Brûlure active pour 2 secondes !"
				}

				if sk.EffectType == "etourdissement" && rand.Intn(100) < 60 {
					enemyStunned = true
					combatLog += " Ennemi étourdi !"
				}
			} else {
				continue
			}
		} else if act == "0" {
			player.CurrentMagie -= 10
			if player.CurrentMagie < 0 {
				player.CurrentMagie = 0
			}
			FledEnemy = enemy
			combatLog = "Vous êtes une peureuse ! Repli tactique en cours..."
			drawArenaFrame(player, enemy, combatLog, 0, "", true)
			time.Sleep(1200 * time.Millisecond)
			return false
		} else {
			continue
		}

		for i := range player.Skills {
			if player.Skills[i].CurrentCD > 0 {
				player.Skills[i].CurrentCD--
			}
		}

		if enemy.CurrentHP <= 0 {
			break
		}

		drawArenaFrame(player, enemy, combatLog, 0, "", true)
		time.Sleep(500 * time.Millisecond)

		if enemyStunned {
			combatLog = fmt.Sprintf("%s est étourdi et rate son tour.", enemy.Name)
			enemyStunned = false
		} else {
			enemyIcon := "👾"
			if enemy.Type == "drone" {
				enemyIcon = "⚡"
			} else if enemy.Type == "cyborg" {
				enemyIcon = "🔻"
			} else if enemy.Type == "mentor" {
				enemyIcon = "👑"
			}

			combatLog = fmt.Sprintf("Riposte de %s !", enemy.Name)
			animateAttack(player, enemy, combatLog, enemyIcon, false)

			dmg := enemy.Attack
			player.CurrentHP -= dmg
			if player.CurrentHP < 0 {
				player.CurrentHP = 0
			}
			combatLog = fmt.Sprintf("Vous subissez %d dégâts.", dmg)

			if (enemy.Type == "mutant" || enemy.Type == "mentor") && rand.Intn(100) < 50 {
				player.IsBurning = true
				player.BurnTimer = time.Now().Add(2 * time.Second)
				combatLog += " 🔥 Vous êtes brûlé par le monstre pendant 2 secondes !"
			}
		}
	}

	drawArenaFrame(player, enemy, combatLog, 0, "", true)

	if enemy.CurrentHP <= 0 {
		fmt.Printf("\n%s🎉 Ennemi éliminé avec succès !%s\n", Green+Bold, Reset)
		player.FightsWon++
		addExperience(player, enemy.ExpReward)

		coins := rand.Intn(20) + 15
		player.Credits += coins
		fmt.Printf("%s🪙 Butin récupéré : +%d Crédits%s\n", Orange, coins, Reset)

		fmt.Print("\n[Appuyez sur Entrée pour continuer]")
		reader.ReadString('\n')
		return true
	}

	fmt.Print("\n[Appuyez sur Entrée pour continuer]")
	reader.ReadString('\n')
	clearScreen()
	return false
}
