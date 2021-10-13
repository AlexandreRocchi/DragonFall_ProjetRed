package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (v *personnage) forgeron() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("                     <((((((\\ ")
	fmt.Println("                     /   S  . } ")
	fmt.Println(" ________            ;--..--._|}")
	fmt.Println("|        |           '--/ --'  )")
	fmt.Println("|________|           | '-'  :'| ")
	fmt.Println("    ||               . -==- .-|")
	fmt.Println("    ||             _--  _.'   /--._")
	fmt.Println("    ||           __.--|       //  _/'--.")
	fmt.Println("    ||         .'-._ ('-----'/ __/      | ")
	fmt.Println("    ||        /  __>|      | '--.       |")
	fmt.Println("    L|       |                          |")
	fmt.Println("___________________________________________")
	fmt.Println("                 Forgeron                  ")
	fmt.Println("___________________________________________")
	fmt.Println("")
	fmt.Println("")
	SlowPrint(Yellow + "Alors mon p'tit! Viens donc dépenser tes pièces dans d'la bonne qualité!\n" + Reset)
	fmt.Println("Fabriquer un chapeau de l'aventurier  (1)")
	fmt.Println("Fabriquer une tunique de l'aventurier (2)")
	fmt.Println("Fabriquer des bottes de l'aventurier  (3)")
	fmt.Println("Retour (10)")
	scan_forgeron := bufio.NewScanner(os.Stdin)
	scan_forgeron.Scan()
	nom_forgeron, _ := strconv.Atoi(scan_forgeron.Text())
	if nom_forgeron == 1 {
		compo_corbeau := 0
		compo_sanglier := 0
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Plume de corbeau" {
				compo_corbeau = compo_corbeau + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()
			}
		}
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Cuir de sanglier" {
				compo_sanglier = compo_sanglier + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()
			}
		}
		if compo_corbeau >= 1 && compo_sanglier >= 1 {
			fmt.Println("Chapeau de l'aventurier obtenu")
			v.addinventory("Chapeau de l'aventurier", 5)
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.forgeron()
		} else {
			fmt.Println("Vous n'avez pas les composants requis pour cet objet")
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.menu()
		}
	}
	if nom_forgeron == 2 {
		compo_loup := 0
		compo_troll := 0
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Fourrure de loup" {
				compo_loup = compo_loup + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()
			}
		}
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Peau de troll" {
				compo_troll = compo_troll + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()

			}
		}
		if compo_loup >= 2 && compo_troll >= 1 {
			fmt.Println("Tunique de l'aventurier obtenue")
			v.addinventory("Tunique de l'aventurier", 5)
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.forgeron()
		} else {
			fmt.Println("Vous n'avez pas les composants requis pour cet objet")
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.menu()
		}
	}
	if nom_forgeron == 10 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
	if nom_forgeron == 3 {
		compo_loup := 0
		compo_sanglier := 0
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Fourrure de loup" {
				compo_loup = compo_loup + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()
			}
		}
		for i := 0; i <= len(v.inventaire)-1; i++ {
			if v.inventaire[i] == "Cuir de sanglier" {
				compo_sanglier = compo_sanglier + 1
				v.inventaire[i] = "Emplacement vide"
				v.removeinventory()
			}
		}
		if compo_loup >= 1 && compo_sanglier >= 1 {
			fmt.Println("Bottes de l'aventurier obtenues")
			v.addinventory("Bottes de l'aventurier", 5)
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.forgeron()
		} else {
			fmt.Println("Vous n'avez pas les composants requis pour cet objet")
			time.Sleep(1 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.forgeron()
		}
	}
}
