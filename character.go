package main

import "fmt"

// Structure du personnage
type Character struct {
	Name            string
	HP              int
	MaxHP           int
	Damage          int
	Stuff bool   // Bloque la salle 1 si 'false'
	Inventory       []Item // Liste des objets ramassés
}

type Item struct {
	Name        string
	Type        string // "Arme", "Armure", "Potion"
	BonusHP     int
	BonusDamage int
}

// Crée un nouveau joueur
func NewCharacter(name string) Character {
	return Character{
		Name:      name,
		HP:        100,
		MaxHP:     100,
		Damage:    10,
		Stuff:     false, // Bloqué au début
		Inventory: []Item{},
	}
}

// Récupère le stuff au vestiaire
func (c *Character) EquipStarterStuff() {
	c.Stuff = true
	c.MaxHP += 20
	c.HP = c.MaxHP
	c.Damage += 5
	fmt.Println("\n✓ Équipement de départ récupéré ! (PV +30, Dégâts +5)")
}