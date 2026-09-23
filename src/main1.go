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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
	// Si le terminal n'efface pas, on peut ajouter un affichage de retours à la ligne de secours :
	// fmt.Println(strings.Repeat("\n", 50))
}

func openInventoryMenu(player *Character, reader *bufio.Reader) {
	for {
		clearScreen()
		fmt.Println(Cyan + "--- 🎒 GESTION DU SAC À DOS  ---" + Reset)
		fmt.Printf(" ❤️ %d/%d  |  ⭐ %d/%d\n", player.CurrentHP, player.MaxHP, player.CurrentMagie, player.MaxMagie)
		fmt.Println(Gray + "------------------------------------------------" + Reset)

		if len(player.Inventory) == 0 {
			fmt.Println(" (Votre inventaire est actuellement vide)")
		} else {
			for i, item := range player.Inventory {
				fmt.Printf(" [%d] 📦 %s\n", i+1, item)
			}
		}

		fmt.Println("\n [Numéro d'objet] Utiliser  |  [0] Quitter")
		fmt.Print("► Choix : ")

		in, _ := reader.ReadString('\n')
		c := strings.TrimSpace(in)

		if c == "0" {
			break
		}

		val, err := strconv.Atoi(c)
		if err == nil && val >= 1 && val <= len(player.Inventory) {
			item := player.Inventory[val-1]
			if item == "🍷 Potion de soin basique" {
				if player.CurrentHP >= player.MaxHP {
					fmt.Println(Yellow + "Vos ❤️ sont déjà au maximum !" + Reset)
				} else {
					player.CurrentHP += 20
					if player.CurrentHP > player.MaxHP {
						player.CurrentHP = player.MaxHP
					}
					removeItem(player, item)
					fmt.Println(Green + "✨ Potion de soin basique utilisée (+20 ❤️) !" + Reset)
				}
			} else if item == "🍷 Potion de soin supérieure" {
				if player.CurrentHP >= player.MaxHP {
					fmt.Println(Yellow + "Vos ❤️ sont déjà au maximum !" + Reset)
				} else {
					player.CurrentHP += 50
					if player.CurrentHP > player.MaxHP {
						player.CurrentHP = player.MaxHP
					}
					removeItem(player, item)
					fmt.Println(Green + "✨ Potion de soin supérieure utilisée (+50 ❤️) !" + Reset)
				}
			} else if item == "🧪 Potion de magie basique" {
				if player.CurrentMagie >= player.MaxMagie {
					fmt.Println(Yellow + "Votre ⭐ est déjà au maximum !" + Reset)
				} else {
					player.CurrentMagie += 20
					if player.CurrentMagie > player.MaxMagie {
						player.CurrentMagie = player.MaxMagie
					}
					removeItem(player, item)
					fmt.Println(Green + "✨ Potion de magie basique utilisée (+20 ⭐) !" + Reset)
				}
			} else if item == "🧪 Potion de magie supérieure" {
				if player.CurrentMagie >= player.MaxMagie {
					fmt.Println(Yellow + "Votre ⭐ est déjà au maximum !" + Reset)
				} else {
					player.CurrentMagie += 40
					if player.CurrentMagie > player.MaxMagie {
						player.CurrentMagie = player.MaxMagie
					}
					removeItem(player, item)
					fmt.Println(Green + "✨ Potion de magie supérieure utilisée (+40 ⭐) !" + Reset)
				}
			} else {
				fmt.Println(Yellow + "Cet objet ne peut pas être consommé directement." + Reset)
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func openMerchant(player *Character, reader *bufio.Reader) {
	for {
		clearScreen()
		fmt.Println(Cyan + "--- 🛒 MARCHAND D'YNOV (STOCKS LIMITÉS) ---" + Reset)
		fmt.Printf(" 🪙 Bourse : %d  |  🎒 Sac : %d/%d\n", player.Credits, len(player.Inventory), player.MaxInventory)
		fmt.Println(Gray + "------------------------------------------------" + Reset)

		fmt.Printf(" [1] 🍷 Potion de soin basique (+20 ❤️) (3 🪙)      - Stock: %d\n", MerchantStock["🍷 Potion de soin basique"])
		fmt.Printf(" [2] 🍷 Potion de soin supérieure (+50 ❤️) (5 🪙)   - Stock: %d\n", MerchantStock["🍷 Potion de soin supérieure"])
		fmt.Printf(" [3] 🧪 Potion de magie basique (+20 ⭐) (3 🪙)      - Stock: %d\n", MerchantStock["🧪 Potion de magie basique"])
		fmt.Printf(" [4] 🧪 Potion de magie supérieure (+40 ⭐) (5 🪙)   - Stock: %d\n", MerchantStock["🧪 Potion de magie supérieure"])
		fmt.Printf(" [5] 🔥 Boule de Feu [Dégâts: 20] (25 🪙)           - Stock: %d\n", MerchantStock["🔥 Boule de Feu"])
		fmt.Printf(" [6] ☕ Jet de Café [Dégâts: 14] (20 🪙)            - Stock: %d\n", MerchantStock["☕ Jet de Café Brûlant"])
		fmt.Printf(" [7] 💻 Ctrl+Alt+Suppr [Dégâts: 30] (30 🪙)         - Stock: %d\n", MerchantStock["💻 Ctrl+Alt+Suppr"])
		fmt.Println(" --- Matériaux d'artisanat ---")
		fmt.Printf(" [8] 🪶 Plume de Corbeau (6 🪙)                   - Stock: %d\n", MerchantStock["🪶 Plume de Corbeau"])
		fmt.Printf(" [9] 🐗 Cuir de Sanglier (6 🪙)                   - Stock: %d\n", MerchantStock["🐗 Cuir de Sanglier"])
		fmt.Printf(" [10] 🐺 Fourrure de Loup (8 🪙)                  - Stock: %d\n", MerchantStock["🐺 Fourrure de Loup"])
		fmt.Printf(" [11] 👹 Peau de Troll (12 🪙)                    - Stock: %d\n", MerchantStock["👹 Peau de Troll"])
		fmt.Println(" [0] Quitter")
		fmt.Print("\n► Votre choix : ")

		in, _ := reader.ReadString('\n')
		c := strings.TrimSpace(in)

		if c == "0" {
			break
		}

		itemBought := ""
		cost := 0
		stockKey := ""

		switch c {
		case "1":
			stockKey = "🍷 Potion de soin basique"
			if MerchantStock[stockKey] > 0 && player.Credits >= 3 && canAddItem(player) {
				cost = 3
				itemBought = stockKey
			}
		case "2":
			stockKey = "🍷 Potion de soin supérieure"
			if MerchantStock[stockKey] > 0 && player.Credits >= 5 && canAddItem(player) {
				cost = 5
				itemBought = stockKey
			}
		case "3":
			stockKey = "🧪 Potion de magie basique"
			if MerchantStock[stockKey] > 0 && player.Credits >= 3 && canAddItem(player) {
				cost = 3
				itemBought = stockKey
			}
		case "4":
			stockKey = "🧪 Potion de magie supérieure"
			if MerchantStock[stockKey] > 0 && player.Credits >= 5 && canAddItem(player) {
				cost = 5
				itemBought = stockKey
			}
		case "5":
			stockKey = "🔥 Boule de Feu"
			if MerchantStock[stockKey] > 0 && player.Credits >= 25 {
				MerchantStock[stockKey]--
				player.Credits -= 25
				player.Skills = append(player.Skills, Skill{Name: "Boule de Feu", Damage: 20, MagieCost: 15, Description: "Attaque de feu", EffectType: "brulure", Icon: "🔥", CooldownMax: 2, CurrentCD: 0})
				fmt.Println(Green + "✨ Sort 'Boule de Feu' mémorisé !" + Reset)
				time.Sleep(1000 * time.Millisecond)
				continue
			}
		case "6":
			stockKey = "☕ Jet de Café Brûlant"
			if MerchantStock[stockKey] > 0 && player.Credits >= 20 {
				MerchantStock[stockKey]--
				player.Credits -= 20
				player.Skills = append(player.Skills, Skill{Name: "Jet de Café Brûlant", Damage: 14, MagieCost: 8, Description: "Brûle l'ennemi avec de la caféine pure", EffectType: "brulure", Icon: "☕", CooldownMax: 1, CurrentCD: 0})
				fmt.Println(Green + "✨ Sort 'Jet de Café Brûlant' mémorisé !" + Reset)
				time.Sleep(1000 * time.Millisecond)
				continue
			}
		case "7":
			stockKey = "💻 Ctrl+Alt+Suppr"
			if MerchantStock[stockKey] > 0 && player.Credits >= 30 {
				MerchantStock[stockKey]--
				player.Credits -= 30
				player.Skills = append(player.Skills, Skill{Name: "Ctrl+Alt+Suppr", Damage: 30, MagieCost: 22, Description: "Redémarre le cerveau du monstre de force", EffectType: "degats", Icon: "💻", CooldownMax: 3, CurrentCD: 0})
				fmt.Println(Green + "✨ Sort 'Ctrl+Alt+Suppr' mémorisé !" + Reset)
				time.Sleep(1000 * time.Millisecond)
				continue
			}
		case "8":
			stockKey = "🪶 Plume de Corbeau"
			if MerchantStock[stockKey] > 0 && player.Credits >= 6 && canAddItem(player) {
				cost = 6
				itemBought = stockKey
			}
		case "9":
			stockKey = "🐗 Cuir de Sanglier"
			if MerchantStock[stockKey] > 0 && player.Credits >= 6 && canAddItem(player) {
				cost = 6
				itemBought = stockKey
			}
		case "10":
			stockKey = "🐺 Fourrure de Loup"
			if MerchantStock[stockKey] > 0 && player.Credits >= 8 && canAddItem(player) {
				cost = 8
				itemBought = stockKey
			}
		case "11":
			stockKey = "👹 Peau de Troll"
			if MerchantStock[stockKey] > 0 && player.Credits >= 12 && canAddItem(player) {
				cost = 12
				itemBought = stockKey
			}
		}

		if itemBought != "" {
			MerchantStock[stockKey]--
			player.Credits -= cost
			player.Inventory = append(player.Inventory, itemBought)
			fmt.Println(Green + "✅ Achat effectué : " + itemBought + " !" + Reset)
		} else if stockKey != "" && MerchantStock[stockKey] <= 0 {
			fmt.Println(Red + "❌ Rupture de stock pour cet article !" + Reset)
		} else if cost > 0 {
			fmt.Println(Red + "❌ Crédits insuffisants ou sac plein !" + Reset)
		} else {
			fmt.Println(Red + "❌ Choix invalide." + Reset)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func open3DPrinter(player *Character, reader *bufio.Reader) {
	for {
		clearScreen()
		fmt.Println(Cyan + "--- 🖨️ IMPRIMANTE 3D D'YNOV (STOCKS LIMITÉS) ---" + Reset)
		fmt.Printf(" 🪙 Bourse : %d\n", player.Credits)
		fmt.Println(Gray + "------------------------------------------------" + Reset)

		fmt.Printf(" [1] 🦄 Corne de licorne (+10 ❤️) (1 Plume + 1 Cuir + 5 🪙)       - Stock: %d\n", PrinterStock["🦄 Corne de licorne"])
		fmt.Printf(" [2] 🐉 Armure en écaille de dragon (+25 ❤️) (2 Fourrures + 1 Peau + 5 🪙) - Stock: %d\n", PrinterStock["🐉 Armure en écaille de dragon"])
		fmt.Printf(" [3] 🥾 Botte de l'avatar (+15 ❤️)  (1 Fourrure + 1 Cuir + 5 🪙)    - Stock: %d\n", PrinterStock["🥾 Botte de l'avatar"])
		fmt.Println(" [0] Quitter")
		fmt.Print("\n► Votre choix : ")

		in, _ := reader.ReadString('\n')
		c := strings.TrimSpace(in)

		if c == "0" {
			break
		}

		itemKey := ""
		switch c {
		case "1":
			itemKey = "🦄 Corne de licorne"
		case "2":
			itemKey = "🐉 Armure en écaille de dragon"
		case "3":
			itemKey = "🥾 Botte de l'avatar"
		}

		if itemKey == "" {
			continue
		}

		if PrinterStock[itemKey] <= 0 {
			fmt.Println(Red + "\n❌ Rupture de stock sur cette impression !" + Reset)
			time.Sleep(1 * time.Second)
			continue
		}

		if player.Credits < 5 {
			fmt.Println(Red + "\n❌ Crédits insuffisants (5 🪙 requis) !" + Reset)
			time.Sleep(1 * time.Second)
			continue
		}

		switch itemKey {
		case "🦄 Corne de licorne":
			if countItem(player, "🪶 Plume de Corbeau") >= 1 && countItem(player, "🐗 Cuir de Sanglier") >= 1 {
				PrinterStock[itemKey]--
				player.Credits -= 5
				removeItem(player, "🪶 Plume de Corbeau")
				removeItem(player, "🐗 Cuir de Sanglier")
				equipItemDirect(player, itemKey)
			} else {
				fmt.Println(Red + "\n❌ Matériaux manquants !" + Reset)
			}
		case "🐉 Armure en écaille de dragon":
			if countItem(player, "🐺 Fourrure de Loup") >= 2 && countItem(player, "👹 Peau de Troll") >= 1 {
				PrinterStock[itemKey]--
				player.Credits -= 5
				removeItem(player, "🐺 Fourrure de Loup")
				removeItem(player, "🐺 Fourrure de Loup")
				removeItem(player, "👹 Peau de Troll")
				equipItemDirect(player, itemKey)
			} else {
				fmt.Println(Red + "\n❌ Matériaux manquants !" + Reset)
			}
		case "🥾 Botte de l'avatar":
			if countItem(player, "🐺 Fourrure de Loup") >= 1 && countItem(player, "🐗 Cuir de Sanglier") >= 1 {
				PrinterStock[itemKey]--
				player.Credits -= 5
				removeItem(player, "🐺 Fourrure de Loup")
				removeItem(player, "🐗 Cuir de Sanglier")
				equipItemDirect(player, itemKey)
			} else {
				fmt.Println(Red + "\n❌ Matériaux manquants !" + Reset)
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func displayInfo(c *Character, reader *bufio.Reader) {
	clearScreen()
	fmt.Println(Cyan + "--- 🌌 FICHE DE PERSONNAGE (YNOV) ---" + Reset)
	fmt.Printf(" Nom : %s%s%s  |  Classe : %s%s%s\n", Bold, c.Name, Reset, Purple, c.Class, Reset)
	fmt.Printf(" Niveau : %s%d%s     |  Crédits : %s%d 🪙%s\n", Yellow, c.Level, Reset, Orange, c.Credits, Reset)
	fmt.Printf(" ❤️ %d/%d  |  ⭐ %d/%d | 🏆 Victoires : %d\n", c.CurrentHP, c.MaxHP, c.CurrentMagie, c.MaxMagie, c.FightsWon)
	displayExpBar(c.Exp, c.MaxExp)
	fmt.Println(Gray + "------------------------------------------------" + Reset)

	fmt.Println(Bold + "🛡️ Équipement Actif :" + Reset)
	fmt.Printf("  Tête  : %s\n", c.Equipment.Head)
	fmt.Printf("  Torse : %s\n", c.Equipment.Body)
	fmt.Printf("  Pieds : %s\n", c.Equipment.Feet)

	fmt.Print("\n[Appuyez sur Entrée pour revenir]")
	reader.ReadString('\n')
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	fmt.Println(Red + Bold + `
                                                                                                        
                                                                                                        
													                                                                                                                                                          
                                                                                                                                                          
                                                  YYYYYYY       YYYYYYYNNNNNNNN        NNNNNNNN     OOOOOOOOO     VVVVVVVV           VVVVVVVV             
                                                  Y:::::Y       Y:::::YN:::::::N       N::::::N   OO:::::::::OO   V::::::V           V::::::V             
                                                  Y:::::Y       Y:::::YN::::::::N      N::::::N OO:::::::::::::OO V::::::V           V::::::V             
                                                  Y::::::Y     Y::::::YN:::::::::N     N::::::NO:::::::OOO:::::::OV::::::V           V::::::V             
                                                  YYY:::::Y   Y:::::YYYN::::::::::N    N::::::NO::::::O   O::::::O V:::::V           V:::::V              
                                                     Y:::::Y Y:::::Y   N:::::::::::N   N::::::NO:::::O     O:::::O  V:::::V         V:::::V        :::::: 
                                                      Y:::::Y:::::Y    N:::::::N::::N  N::::::NO:::::O     O:::::O   V:::::V       V:::::V         :::::: 
                                                       Y:::::::::Y     N::::::N N::::N N::::::NO:::::O     O:::::O    V:::::V     V:::::V          :::::: 
                                                        Y:::::::Y      N::::::N  N::::N:::::::NO:::::O     O:::::O     V:::::V   V:::::V                  
                                                         Y:::::Y       N::::::N   N:::::::::::NO:::::O     O:::::O      V:::::V V:::::V                   
                                                         Y:::::Y       N::::::N    N::::::::::NO:::::O     O:::::O       V:::::V:::::V                    
                                                         Y:::::Y       N::::::N     N:::::::::NO::::::O   O::::::O        V:::::::::V              :::::: 
                                                         Y:::::Y       N::::::N      N::::::::NO:::::::OOO:::::::O         V:::::::V               :::::: 
                                                      YYYY:::::YYYY    N::::::N       N:::::::N OO:::::::::::::OO           V:::::V                :::::: 
                                                      Y:::::::::::Y    N::::::N        N::::::N   OO:::::::::OO              V:::V                        
                                                      YYYYYYYYYYYYY    NNNNNNNN         NNNNNNN     OOOOOOOOO                 VVV                         
                                                                                                                                                          
                                                                                                                                                          
                                                                                                                                                          
                                                                                                                                                          
    ` + Reset)

	fmt.Println(Red + Bold + `
LLLLLLLLLLL                  OOOOOOOOO        SSSSSSSSSSSSSSS TTTTTTTTTTTTTTTTTTTTTTT             CCCCCCCCCCCCC     OOOOOOOOO     LLLLLLLLLLL                  OOOOOOOOO     NNNNNNNN        NNNNNNNNYYYYYYY       YYYYYYY
L:::::::::L                OO:::::::::OO    SS:::::::::::::::ST:::::::::::::::::::::T          CCC::::::::::::C   OO:::::::::OO   L:::::::::L                OO:::::::::OO   N:::::::N       N::::::NY:::::Y       Y:::::Y
L:::::::::L              OO:::::::::::::OO S:::::SSSSSS::::::ST:::::::::::::::::::::T        CC:::::::::::::::C OO:::::::::::::OO L:::::::::L              OO:::::::::::::OO N::::::::N      N::::::NY:::::Y       Y:::::Y
LL:::::::LL             O:::::::OOO:::::::OS:::::S     SSSSSSST:::::TT:::::::TT:::::T       C:::::CCCCCCCC::::CO:::::::OOO:::::::OLL:::::::LL             O:::::::OOO:::::::ON:::::::::N     N::::::NY::::::Y     Y::::::Y
  L:::::L               O::::::O   O::::::OS:::::S            TTTTTT  T:::::T  TTTTTT      C:::::C       CCCCCCO::::::O   O::::::O  L:::::L               O::::::O   O::::::ON::::::::::N    N::::::NYYY:::::Y   Y:::::YYY
  L:::::L               O:::::O     O:::::OS:::::S                    T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON:::::::::::N   N::::::N   Y:::::Y Y:::::Y   
  L:::::L               O:::::O     O:::::O S::::SSSS                 T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON:::::::N::::N  N::::::N    Y:::::Y:::::Y    
  L:::::L               O:::::O     O:::::O  SS::::::SSSSS            T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON::::::N N::::N N::::::N     Y:::::::::Y     
  L:::::L               O:::::O     O:::::O    SSS::::::::SS          T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON::::::N  N::::N:::::::N      Y:::::::Y      
  L:::::L               O:::::O     O:::::O       SSSSSS::::S         T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON::::::N   N:::::::::::N       Y:::::Y       
  L:::::L               O:::::O     O:::::O            S:::::S        T:::::T             C:::::C              O:::::O     O:::::O  L:::::L               O:::::O     O:::::ON::::::N    N::::::::::N       Y:::::Y       
  L:::::L         LLLLLLO::::::O   O::::::O            S:::::S        T:::::T              C:::::C       CCCCCCO::::::O   O::::::O  L:::::L         LLLLLLO::::::O   O::::::ON::::::N     N:::::::::N       Y:::::Y       
LL:::::::LLLLLLLLL:::::LO:::::::OOO:::::::OSSSSSSS     S:::::S      TT:::::::TT             C:::::CCCCCCCC::::CO:::::::OOO:::::::OLL:::::::LLLLLLLLL:::::LO:::::::OOO:::::::ON::::::N      N::::::::N       Y:::::Y       
L::::::::::::::::::::::L OO:::::::::::::OO S::::::SSSSSS:::::S      T:::::::::T              CC:::::::::::::::C OO:::::::::::::OO L::::::::::::::::::::::L OO:::::::::::::OO N::::::N       N:::::::N    YYYY:::::YYYY    
L::::::::::::::::::::::L   OO:::::::::OO   S:::::::::::::::SS       T:::::::::T                CCC::::::::::::C   OO:::::::::OO   L::::::::::::::::::::::L   OO:::::::::OO   N::::::N        N::::::N    Y:::::::::::Y    
LLLLLLLLLLLLLLLLLLLLLLLL     OOOOOOOOO      SSSSSSSSSSSSSSS         TTTTTTTTTTT                   CCCCCCCCCCCCC     OOOOOOOOO     LLLLLLLLLLLLLLLLLLLLLLLL     OOOOOOOOO     NNNNNNNN         NNNNNNN    YYYYYYYYYYYYY    
                                                                                                                                                                                                                          


	` + Reset)

	fmt.Println(`

 n                                                                 :.
 E%                                                                :"5
z  %                                                              :" "
K   ":                                                           z   R
?     %.                                                       :^    J
 ".    ^s                                                     f     :~
  '+.    #L                                                 z"    .*
    '+     %L                                             z"    .~
      ":    '%.                                         .#     +
        ":    ^%.                                     .#"    +"
          #:    "n                                  .+"   .z"
            #:    ":                               z"    +"
              %:   "*L                           z"    z"
                *:   ^*L                       z*   .+"
                  "s   ^*L                   z#   .*"
                    #s   ^%L               z#   .*"
                      #s   ^%L           z#   .r"
                        #s   ^%.       u#   .r"
                          #i   '%.   u#   .@"
                            #s   ^%u#   .@"
                              #s x#   .*"
                               x#"  .@%.
                             x#"  .d"  "%.
                           xf~  .r" #s   "%.
                     u   x*"  .r"     #s   "%.  x.
                     %Mu*"  x*"         #m.  "%zX"
                     :R(h x*              "h..*dN.
                   u@NM5e#>                 7?dMRMh.
                 z$@M@$#"#"                 *""*@MM$hL
               u@@MM8*                          "*$M@Mh.
             z$RRM8F"                             "N8@M$bL
            5"RM$#                                  "R88f)R
            'h.$"                                     #$x*


` + Reset)

	fmt.Print(Purple + `--     _  _   __  ____  ____  ____    __ _   __   _  _    ____  ____    _  _  ____  ____   __   ____    _   
--    / )( \ /  \(_  _)(  _ \(  __)  (  ( \ /  \ ( \/ )  (    \(  __)  / )( \(  __)(  _ \ /  \ / ___)  (_)  
--    \ \/ /(  O ) )(   )   / ) _)   /    /(  O )/ \/ \   ) D ( ) _)   ) __ ( ) _)  )   /(  O )\___ \   _   
--     \__/  \__/ (__) (__\_)(____)  \_)__) \__/ \_)(_/  (____/(____)  \_)(_/(____)(__\_) \__/ (____/  (_)   `)
	nameIn, _ := reader.ReadString('\n')
	rawName := strings.TrimSpace(nameIn)

	var name string
	if len(rawName) > 0 {
		runes := []rune(rawName)
		name = strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
	} else {
		name = "Survivant"
	}
	fmt.Print(Reset)

	var class string
	for {
		fmt.Println(Cyan + "\nChoisissez votre classe :" + Reset)
		fmt.Println(" [1] 👤 Humain (❤️ 100 | ⭐ 50)  - (Moyen) Équilibré & Polyvalent")
		fmt.Println(" [2] 🧝 Elfe   (❤️ 80  | ⭐ 80)  - (Difficile) Agile & Mage Puissant")
		fmt.Println(" [3] 🛡️ Nain   (❤️ 120 | ⭐ 30)  - (Facile) Solide & Gros Tank")
		fmt.Print(Cyan + "► Choix (1, 2 ou 3) : " + Reset)

		cIn, _ := reader.ReadString('\n')
		cClean := strings.TrimSpace(cIn)

		if cClean == "1" {
			class = "Humain"
			break
		} else if cClean == "2" {
			class = "Elfe"
			break
		} else if cClean == "3" {
			class = "Nain"
			break
		} else {
			fmt.Println(Red + "❌ Choix invalide. Entrez 1, 2 ou 3." + Reset)
		}
	}

	p := initCharacter(name, class)
	player := &p

	for {
		if player.CurrentHP <= 0 {
			fmt.Println(Red + "\n💀 Vous avez succombé ! Réanimation d'urgence à 50% des ❤️..." + Reset)
			player.CurrentHP = player.MaxHP / 2
			time.Sleep(1200 * time.Millisecond)
		}

		clearScreen()
		fmt.Println(Cyan + "--- 🕹️ MENU PRINCIPAL (YNOV) ---" + Reset)
		fmt.Printf(" Héros : %s%s (%s)%s | Niv: %d | ❤️: %d/%d | 🏆 Victoires : %d\n", Bold, player.Name, player.Class, Reset, player.Level, player.CurrentHP, player.MaxHP, player.FightsWon)
		fmt.Println(Gray + "--------------------------------------------------------" + Reset)
		fmt.Println(" [1] 🗺️  Explorer Ynov (Grille 15x15)")
		fmt.Println(" [2] ⚔️  Lancer un Combat Aléatoire")
		fmt.Println(" [3] 👤 Fiche de Personnage")
		fmt.Println(" [4] 🛒 Marchand (Potions, Sorts & Matériaux)")
		fmt.Println(" [5] 🖨️ Imprimante 3D (Équipements)")
		fmt.Println(" [6] 🎒 Sac à dos (Objets)")

		// Option pour les 3 mentors débloquée à partir de 3 victoires
		if player.FightsWon >= 3 {
			fmt.Println(Green + " [7] 🎓 Affronter les 3 Mentors (BOSS FINAL) [DÉBLOQUÉ !]" + Reset)
		} else {
			fmt.Println(Gray + " [7] 🔒 Affronter les 3 Mentors (Requis : 3 victoires)" + Reset)
		}
		fmt.Println(" [8] 🎒 Qui sont-ils (Easter egg)")

		fmt.Println(" [0] 🚪 Quitter")
		fmt.Println(Gray + "--------------------------------------------------------" + Reset)
		fmt.Print("► Votre choix : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			exploreMap(player, reader)
		case "2":
			enemy := generateRandomMonster(player.Level, player.Exp)
			runFight(player, &enemy, reader)
		case "3":
			displayInfo(player, reader)
		case "4":
			openMerchant(player, reader)
		case "5":
			open3DPrinter(player, reader)
		case "6":
			openInventoryMenu(player, reader)
		case "7":
			if player.FightsWon >= 3 {
				clearScreen()
				fmt.Println(Red + Bold + "⚡ ALERTE ROUGE : AFFRONTEMENT SUPRÊME CONTRE LES 3 MENTORS ! ⚡" + Reset)
				fmt.Println("Les 3 Mentors d'Ynov vous font face pour l'examen final de votre vie !")
				fmt.Print("\n[Appuyez sur Entrée pour lancer le combat final]")
				reader.ReadString('\n')

				// Création du Boss des Mentors (qui va chercher ascii.txt)
				bossMentors := Enemy{
					Name:       "Les 3 Mentors d'Ynov",
					MaxHP:      250,
					CurrentHP:  250,
					Attack:     20,
					Initiative: 20,
					ExpReward:  1000,
					Type:       "mentor", // Ceci charge le fichier ascii.txt
				}

				wonBoss := runFight(player, &bossMentors, reader)

				if wonBoss {
					clearScreen()
					fmt.Println(Green + Bold + "╔════════════════════════════════════════════════════════════╗")
					fmt.Println("║    🎉 VOUS AVEZ VAINCU LES MENTORS ET GAGNÉ LE JEU ! 🎉   ║")
					fmt.Println("╚════════════════════════════════════════════════════════════╝" + Reset)
					time.Sleep(800 * time.Millisecond) // Petite pause pour lire le message avant de couper
					fmt.Printf(Blue + `
 _____ __ _ _      _ _        _   _                            
|  ___/_/| (_) ___(_) |_ __ _| |_(_) ___  _ __  ___            
| |_ / _ \ | |/ __| | __/ _  | __| |/ _ \|  _ \/ __|           
|  _|  __/ | | (__| | || (_| | |_| | (_) | | | \__ \           
|_|  \___|_|_|\___|_|\__\__,_|\__|_|\___/|_| |_|___/     _     
 _ __ ___   __ _  __ _(_)___| |_ _ __ __ _| | ___  ___  | |    
| '_   _ \ / _  |/ _  | / __| __| '__/ _  | |/ _ \/ __| | |    
| | | | | | (_| | (_| | \__ \ |_| | | (_| | |  __/\__ \ |_|    
|_| |_|_|_|\__,_|\__, |_|___/\__|_|  \__,_|_|\___||___/ (_)    
\ \   / /__  _   |___/   _ __ ___ _ __   __ _ _ __| |_ ___ ____
 \ \ / / _ \| | | / __| | '__/ _ \ '_ \ / _  |  __| __/ _ \_  /
  \ V / (_) | |_| \__ \ | | |  __/ |_) | (_| | |  | ||  __// / 
   \_/ \___/ \__,_|___/ |_|  \___| .__/ \__,_|_|   \__\___/___|
                                 |_|  ` + Reset)
					fmt.Println(`                          
  __ ___   _____  ___  __   _____ | |_ _ __ ___                
 / _  \ \ / / _ \/ __| \ \ / / _ \| __| '__/ _ \               
| (_| |\ V /  __/ (__   \ V / (_) | |_| | |  __/               
 \__,_|_\_/ \___|\//\|   \_/ \___/ \__|_|  \___|               
  __| (_)_ __ | ||/_\| _ __ ___   ___                          
 / _  | | '_ \| |/ _ \| '_   _ \ / _ \                         
| (_| | | |_) | | (_) | | | | | |  __/                         
 \__,_|_| .__/|_|\___///\ |_| |_|\___|                         
 ___ _  |_|_ __  _ __|/_\|_ __ ___   ___                       
/ __| | | | '_ \| '__/ _ \ '_   _ \ / _ \                      
\__ \ |_| | |_) | | |  __/ | | | | |  __/_                     
|___/\__,_| .__/|_|  \___|_| |_| |_|\___(_)                    
          |_|                                              ` + Reset)
					time.Sleep(800 * time.Millisecond) // Petite pause pour lire le message avant de couper

					fmt.Println(Red + `
 _                                                __   __                 
| |    ___    ___ __ _ _ __ ___  _ __  _   _ ___  \ \ / / __   _____   __ 
| |   / _ \  / __/ _  | '_   _ \|  _ \| | | / __|  \ V /  _ \ / _ \ \ / / 
| |__|  __/ | (_| (_| | | | | | | |_) | |_| \__ \   | || | | | (_) \ V /  
|_____\___| _\___\__,_|_| |_| |_| .__/ \__,_|___/   |_||_| |_|\___/ \_/   
  ___| |_  | | ___  ___   _ __ _|_|  ___ _ __ | |_ ___  _ __ ___          
 / _ \ __| | |/ _ \/ __| | '_   _ \ / _ \ '_ \| __/ _ \| '__/ __|         
|  __/ |_  | |  __/\__ \ | | | | | |  __/ | | | || (_) | |  \__ \         
 \___|\__| |_|\___||___/ |_| |_| |_|\___|_| |_|\__\___/|_|  |___/_        
 ___( |_)_ __   ___| (_)_ __   ___    __| | _____   ____ _ _ __ | |_      
/ __|/| | '_ \ / __| | | '_ \ / _ \  / _  |/ _ \ \ / / _  |  _ \| __|     
\__ \ | | | | | (__| | | | | |  __/ | (_| |  __/\ V / (_| | | | | |_      
|___/ |_|_| |_|\___|_|_|_| |_|\___|  \__,_|\___| \_/ \__,_|_| |_|\__|   _ 
__   _____ | |_ _ __ ___   _ __  _   _(_)___ ___  __ _ _ __   ___ ___  | |
\ \ / / _ \| __| '__/ _ \ | '_ \| | | | / __/ __|/ _  |  _ \ / __/ _ \ | |
 \ V / (_) | |_| | |  __/ | |_) | |_| | \__ \__ \ (_| | | | | (_|  __/ |_|
  \_/ \___/ \__|_|  \___| | .__/ \__,_|_|___/___/\__,_|_| |_|\___\___| (")  ` + Reset)
					time.Sleep(800 * time.Millisecond) // Petite pause pour lire le message avant de couper
					fmt.Print("\n[Appuyez sur Entrée pour quitter]")
					reader.ReadString('\n')
					return
				}
			} else {
				fmt.Println(Red + "❌ Vous devez avoir au moins 3 victoires pour affronter les mentors !" + Reset)
				time.Sleep(1500 * time.Millisecond)
			}

		case "8":
			showCreators()

			// Appel de la fonction qui affiche les prénoms
		case "0":
			clearScreen() // <--- Nettoie l'écran juste avant de fermer
			fmt.Println(Cyan + "Fermeture du système YNOV..." + Reset)
			time.Sleep(800 * time.Millisecond) // Petite pause pour lire le message avant de couper
			clearScreen()                      // <--- Optionnel : un dernier clear pour tout effacer proprement
			return
		}
	}
}
func showCreators() {
	clearScreen() // Nettoie l'écran proprement

	fmt.Println(Cyan + Bold + "========================================" + Reset)
	fmt.Println(Green + Bold + "         QUI SONT-ILS ?        " + Reset)
	fmt.Println(Cyan + Bold + "========================================" + Reset)

	fmt.Println("\nLes deux personnages dans les parties 2 et 3 sont :\n")

	// Remplacez "Prénom 1" et "Prénom 2" par les vrais prénoms !
	fmt.Println(Yellow + `     
    _    ____  ____    _    
   / \  | __ )| __ )  / \   
  / _ \ |  _ \|  _ \ / _ \  
 / ___ \| |_) | |_) / ___ \ 
/_/   \_\____/|____/_/   \_\ 
	` + Reset)
	fmt.Println(Purple + `
 ____  ____ ___ _____ _     ____  _____ ____   ____ 
/ ___||  _ \_ _| ____| |   | __ )| ____|  _ \ / ___|
\___ \| |_) | ||  _| | |   |  _ \|  _| | |_) | |  _ 
 ___) |  __/| || |___| |___| |_) | |___|  _ <| |_| |
|____/|_|  |___|_____|_____|____/|_____|_| \_\\____|
	` + Reset + "\n")

	fmt.Println(Cyan + "========================================" + Reset)
	fmt.Println("Appuyez sur Entrée pour revenir au menu...")

	// Attend que le joueur appuie sur Entrée pour repartir
	var input string
	fmt.Scanln(&input)
}
