package main

import "time"

const (
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	Red      = "\033[38;5;196m"
	Green    = "\033[38;5;46m"
	Yellow   = "\033[38;5;226m"
	Blue     = "\033[38;5;33m"
	Cyan     = "\033[38;5;51m"
	Orange   = "\033[38;5;208m"
	Purple   = "\033[38;5;141m"
	Gray     = "\033[38;5;242m"
	DarkGray = "\033[38;5;236m"
)

type Equipment struct {
	Head string
	Body string
	Feet string
}

type Skill struct {
	Name        string
	Damage      int
	MagieCost   int
	Description string
	EffectType  string
	Icon        string
	CooldownMax int
	CurrentCD   int
}

type Character struct {
	Name            string
	Class           string
	Level           int
	Exp             int
	MaxExp          int
	MaxHP           int
	CurrentHP       int
	MaxMagie        int
	CurrentMagie    int
	Initiative      int
	Credits         int
	Inventory       []string
	MaxInventory    int
	Equipment       Equipment
	Skills          []Skill
	LastExploreTime time.Time
	IsBurning       bool
	BurnTimer       time.Time
	FightsWon       int
}

type Enemy struct {
	Name       string
	MaxHP      int
	CurrentHP  int
	Attack     int
	Initiative int
	ExpReward  int
	Type       string
	IsBurning  bool
	BurnTimer  time.Time
}

var FledEnemy *Enemy = nil
var StatusMsg string = ""

var MerchantStock = map[string]int{
	"🧪 Potion de soin basique":     5,
	"🧪 Potion de soin supérieure":  5,
	"🧪 Potion de magie basique":    5,
	"🧪 Potion de magie supérieure": 5,
	"🔥 Boule de Feu":               1,
	"☕ Jet de Café Brûlant":        1,
	"💻 Ctrl+Alt+Suppr":             1,
	"🪶 Plume de Corbeau":           4,
	"🐗 Cuir de Sanglier":           4,
	"🐺 Fourrure de Loup":           3,
	"👹 Peau de Troll":              2,
}

var PrinterStock = map[string]int{
	"🦄 Corne de licorne":            2,
	"🐉 Armure en écaille de dragon": 1,
	"🥾 Botte de l'avatar":           2,
}
