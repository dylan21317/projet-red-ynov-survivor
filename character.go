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