package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (v *personnage) addinventory(item string, price int) {
	if len(v.inventaire) == v.inventairemax {
		fmt.Println("Vous n'avez plus de place")
		time.Sleep(1 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
	if v.argent > price {
		v.inventaire = append(v.inventaire, item)
	} else {
		fmt.Println("Vous n'avez pas assez pas d'argent")
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
}

func (v *personnage) removeinventory() {
	var remove []string
	for i := 0; i <= len(v.inventaire)-1; i++ {
		if v.inventaire[i] == "Emplacement vide" {
			v.inventaire[i] = v.inventaire[len(v.inventaire)-1]
			v.inventaire[len(v.inventaire)-1] = "Emplacement vide"
		}
	}
	for i := 0; i <= len(v.inventaire)-2; i++ {
		remove = append(remove, v.inventaire[i])
	}
	v.inventaire = remove
}
func (v *personnage) removesorts() {
	var remove []string
	for i := 0; i <= len(v.sorts)-1; i++ {
		if v.sorts[i] == "Emplacement vide" {
			v.sorts[i] = v.sorts[len(v.sorts)-1]
			v.sorts[len(v.sorts)-1] = "Emplacement vide"
		}
	}
	for i := 0; i <= len(v.sorts)-2; i++ {
		remove = append(remove, v.sorts[i])
	}
	v.sorts = remove
}

func (v *personnage) marchand() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("      _______________")
	fmt.Println("     |       B       |")
	fmt.Println("     |               |")
	fmt.Println("  _______________________")
	fmt.Println("  ===========.===========")
	fmt.Println("   |  ~~~~~     ~~~~~  | ")
	fmt.Println("   |  | O |     | O |  | ")
	fmt.Println("   W   ---  /    ---   W")
	fmt.Println("   |.      |o o|      .| ")
	fmt.Println("   |     #########     |")
	fmt.Println("   |    ## ----- ##    |")
	fmt.Println("   _____________________")
	fmt.Println("___________________________________________")
	fmt.Println("                 Marchand                  ")
	fmt.Println("___________________________________________")
	fmt.Println("")
	fmt.Println("")
	SlowPrint(Yellow + "Salut toi!  Achète moi un truc!\n" + Reset)
	fmt.Println("Acheter une potion de soin pour 3 pièces d'or (4)")
	fmt.Println("Acheter le sort de boule de feu pour 25 pièces d'or (5)")
	fmt.Println("Acheter une potion de poison pour 6 pièces d'or (3)")
	fmt.Println("Acheter une fourrure de Loup pour 4 pièces d'or (6)")
	fmt.Println("Acheter une peau de troll pour 7 pièces d'or (7)")
	fmt.Println("Acheter une cuir de sanglier pour 3 pièces d'or (8)")
	fmt.Println("Acheter une plume de corbeau pour 1 pièces d'or (9)")
	fmt.Println("Acheter une sacoche (plus trois places dans l'inventaire) pour 30 pièces d'or (11)")
	fmt.Println("Retour (10)")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	nom_marchand, _ := strconv.Atoi(scanner3.Text())
	if nom_marchand == 5 {
		v.spellbook("Boule de feu", 25)
		v.argent = v.argent - 25
		fmt.Print("Boule de feu achetée")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 4 {
		v.addinventory("Potion de soin", 3)
		v.argent = v.argent - 3
		fmt.Print("Potion de soin achetée")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 3 {
		v.addinventory("Potion de poison", 6)
		v.argent = v.argent - 6
		fmt.Print("Potion de poison achetée")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 6 {
		v.addinventory("Fourrure de loup", 4)
		v.argent = v.argent - 4
		fmt.Print("Fourrure de loup achetée")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 7 {
		v.addinventory("Peau de troll", 7)
		v.argent = v.argent - 7
		fmt.Print("Peau de troll achetée")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 8 {
		v.addinventory("Cuir de sanglier", 3)
		v.argent = v.argent - 3
		fmt.Print("Cuir de sanglier acheté")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 9 {
		v.addinventory("Plume de corbeau", 1)
		v.argent = v.argent - 1
		fmt.Print("Plume de corbeau acheté")
		time.Sleep(1000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.marchand()
	} else if nom_marchand == 10 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	} else if nom_marchand == 11 {
		if v.inventairemax == 19 {
			fmt.Println("Vous avez atteint la taille maximale")
			time.Sleep(1200 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.marchand()
		} else {
			v.argent = v.argent - 30
			fmt.Println("Trois emplacements d'inventaire ajoutés")
			v.inventairemax = v.inventairemax + 3
			time.Sleep(1200 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.marchand()
		}
	}
}
