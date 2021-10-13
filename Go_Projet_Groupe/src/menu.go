package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (p *personnage) menu() string {
	var retour int
	df := Red + "Ɗ Ʀ ƛ Ɠ Ơ Ɲ  Ƒ ƛ Լ Լ" + Reset
	fmt.Println("")
	fmt.Println("")
	fmt.Println("             _____                                                 ")
	fmt.Println("             \\   /                                                 ")
	fmt.Println("             |   |                                                 ")
	fmt.Println(".__.         |   |_____________________________________________    ")
	fmt.Println("|  |_________|   |                                              \\  ")
	fmt.Println("|  |         |   |_______", df, "___________________\\")
	fmt.Println("|  |_________|   |                                                /")
	fmt.Println("|__|         |   |_____________________________________________ /  ")
	fmt.Println("             |   |                                                 ")
	fmt.Println("             |   |                                                 ")
	fmt.Println("             /___\\                                                 ")
	fmt.Println("_______________________________________________")
	fmt.Println("Profil du personnage (1)   ")
	fmt.Println("_______________________________________________")
	fmt.Println("Inventaire (2)     ")
	fmt.Println("_______________________________________________")
	fmt.Println("Aller voir Barry le marchand (3)     ")
	fmt.Println("______________________________________________")
	fmt.Println("Allez voir Smith le forgeron (4)")
	fmt.Println("______________________________________________")
	fmt.Println("Affronter le dragon (5)")
	fmt.Println("______________________________________________")
	fmt.Println("Quitter le jeu (6)")
	fmt.Println("______________________________________________")
	fmt.Println("Accéder à l'arbre des compétences (7)")
	fmt.Println("______________________________________________")
	fmt.Println("Développé par Nathan et Alex")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nom, _ := strconv.Atoi(scanner.Text())
	if nom == 6 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		os.Exit(0)
	}
	if nom == 7 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.Arbre()
	}
	if nom == 2 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.accesInventory()
		fmt.Println("Utiliser une potion de soin (6)")
		fmt.Println("")
		fmt.Println("Utiliser une potion de poison (7)")
		fmt.Println("")
		fmt.Println("Equiper équipement (8)")
		fmt.Println("")
		fmt.Println("Retour (10)")
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		retour, _ = strconv.Atoi(scanner2.Text())
		if retour == 6 {
			p.takePot()
			time.Sleep(3 * time.Second)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			p.menu()
		}
		if retour == 10 {
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			p.menu()
		}
		if retour == 7 {
			p.poisonPot()
			time.Sleep(2 * time.Second)
			p.Wasted()
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			p.menu()
		}
		if retour == 8 {
			p.equiper()
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			p.menu()
		}
	}
	if nom == 5 {
		return ""
	}
	if nom == 4 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.loadingforgeron()
	}
	if nom == 1 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.DisplayInfo()
		fmt.Println("Retour (10)")
		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		retour, _ = strconv.Atoi(scanner2.Text())
		if retour == 10 {
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			p.menu()
		}
	} else if nom == 3 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.loadingmarchand()
	}
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.menu()
	return ""
}
