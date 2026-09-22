package main

import (
	"fmt"
	"strings"
	"time"
)

func initCharacter(name string, class string) Character {
	maxHP, maxMagie, initVal := 100, 50, 10

	if class == "Elfe" {
		maxHP, maxMagie, initVal = 80, 80, 15
	} else if class == "Nain" {
		maxHP, maxMagie, initVal = 120, 30, 5
	}

	skills := []Skill{
		{Name: "Coup de poing", Damage: 8, MagieCost: 0, Description: "Attaque physique de base", EffectType: "degats", Icon: "👊", CooldownMax: 0, CurrentCD: 0},
		{Name: "Onde Étourdissante", Damage: 5, MagieCost: 10, Description: "Chance d'étourdir l'ennemi", EffectType: "etourdissement", Icon: "⚡", CooldownMax: 2, CurrentCD: 0},
	}

	return Character{
		Name:            name,
		Class:           class,
		Level:           1,
		Exp:             0,
		MaxExp:          100,
		MaxHP:           maxHP,
		CurrentHP:       maxHP,
		MaxMagie:        maxMagie,
		CurrentMagie:    maxMagie,
		Initiative:      initVal,
		Credits:         100,
		Inventory:       []string{"🍷 Potion de soin basique", "🍷 Potion de soin superieur", "🍷 Potion de magie basique"},
		MaxInventory:    15,
		Equipment:       Equipment{Head: "Aucun", Body: "Aucun", Feet: "Aucun"},
		Skills:          skills,
		LastExploreTime: time.Time{},
		FightsWon:       0,
	}
}

func updateMaxHP(c *Character) {
	baseHP := 100
	if c.Class == "Elfe" {
		baseHP = 80
	} else if c.Class == "Nain" {
		baseHP = 120
	}

	bonus := 0
	if c.Equipment.Head == "🦄 Corne de licorne" {
		bonus += 10
	}
	if c.Equipment.Body == "🐉 Armure en écaille de dragon" {
		bonus += 25
	}
	if c.Equipment.Feet == "🥾 Botte de l'avatar" {
		bonus += 15
	}

	c.MaxHP = baseHP + bonus
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}
}

func displayExpBar(currentExp, maxExp int) {
	barWidth := 25
	progress := int(float64(currentExp) / float64(maxExp) * float64(barWidth))
	if progress > barWidth {
		progress = barWidth
	}
	remaining := barWidth - progress

	filledStr := strings.Repeat("█", progress)
	emptyStr := strings.Repeat("░", remaining)
	_ = emptyStr

	fmt.Printf("\n 📊 Progression EXP : [%s%s%s] %d / %d XP\n", Green, filledStr, Reset, currentExp, maxExp)
}

func addExperience(c *Character, expGained int) {
	c.Exp += expGained
	fmt.Printf("\n%s✨ Gain d'expérience : +%d XP%s\n", Yellow+Bold, expGained, Reset)
	displayExpBar(c.Exp, c.MaxExp)

	for c.Exp >= c.MaxExp {
		c.Exp -= c.MaxExp
		c.Level++
		c.MaxExp = int(float64(c.MaxExp) * 1.5)
		c.MaxHP += 15
		c.CurrentHP = c.MaxHP
		c.MaxMagie += 10
		c.CurrentMagie = c.MaxMagie
		c.Initiative += 2

		restockShops()

		fmt.Printf("%s🎉 MONTÉE DE NIVEAU ! Vous êtes maintenant niveau %d ! (Stocks réapprovisionnés)%s\n", Green+Bold, c.Level, Reset)
		displayExpBar(c.Exp, c.MaxExp)
	}
}

func restockShops() {
	MerchantStock["🧪 Potion de soin basique"] += 3
	MerchantStock["🧪 Potion de soin supérieure"] += 3
	MerchantStock["🧪 Potion de magie basique"] += 3
	MerchantStock["🧪 Potion de magie supérieure"] += 3
	MerchantStock["🪶 Plume de Corbeau"] += 2
	MerchantStock["🐗 Cuir de Sanglier"] += 2
	MerchantStock["🐺 Fourrure de Loup"] += 2
	MerchantStock["👹 Peau de Troll"] += 1

	PrinterStock["🦄 Corne de licorne"] += 1
	PrinterStock["🥾 Botte de l'avatar"] += 1
}

func canAddItem(c *Character) bool {
	return len(c.Inventory) < c.MaxInventory
}

func countItem(c *Character, itemName string) int {
	count := 0
	for _, item := range c.Inventory {
		if item == itemName {
			count++
		}
	}
	return count
}

func removeItem(c *Character, itemName string) bool {
	for i, item := range c.Inventory {
		if item == itemName {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			return true
		}
	}
	return false
}

func equipItemDirect(player *Character, itemName string) {
	switch itemName {
	case "🦄 Corne de licorne":
		if player.Equipment.Head != "Aucun" {
			player.Inventory = append(player.Inventory, player.Equipment.Head)
		}
		player.Equipment.Head = itemName
		fmt.Println(Green + "🦄 'Corne de licorne' équipée (+10 ❤️ Max) !" + Reset)

	case "🐉 Armure en écaille de dragon":
		if player.Equipment.Body != "Aucun" {
			player.Inventory = append(player.Inventory, player.Equipment.Body)
		}
		player.Equipment.Body = itemName
		fmt.Println(Green + "🐉 'Armure en écaille de dragon' équipée (+25 ❤️ Max) !" + Reset)

	case "🥾 Botte de l'avatar":
		if player.Equipment.Feet != "Aucun" {
			player.Inventory = append(player.Inventory, player.Equipment.Feet)
		}
		player.Equipment.Feet = itemName
		fmt.Println(Green + "🥾 'Botte de l'avatar' équipée (+15 ❤️ Max) !" + Reset)
	}
	updateMaxHP(player)
}
