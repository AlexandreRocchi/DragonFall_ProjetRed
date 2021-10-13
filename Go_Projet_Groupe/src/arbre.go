package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func (v *personnage) Arbre() {
	fmt.Println(`
    
    
                                                   |----------Attaque->+8-|
                     |----------Hp max->+30--------|                      |                        
    Maximator--------|                               |                      |
                     |----------Argent max->+50----|                      |
                                                   |----------Magie->+5---|    


    
                                     `)
	fmt.Println("Augmenter les Hp maximums pour 20XP (1)")
	fmt.Println("Augmenter la taille de la bourse pour 20XP (2)")
	fmt.Println("Augmenter l'attaque pour 30XP (3)")
	fmt.Println("Augmenter la magie pour 40XP (4)")
	fmt.Println("Retour (10)")
	scanner3 := bufio.NewScanner(os.Stdin)
	scanner3.Scan()
	nom_arbre, _ := strconv.Atoi(scanner3.Text())
	if nom_arbre == 1 {
		if v.xp >= 20 {
			v.hpmax += 30
			v.xp -= 20
			fmt.Println("Les Hp maximums ont été augmenté de 30!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		} else {
			fmt.Println("Vous n'avez pas assez d'XP!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		}
	}
	if nom_arbre == 2 {
		if v.xp >= 20 {
			v.argentmax += 50
			v.xp -= 20
			fmt.Println("La bourse a été augmentée de 50!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		} else {
			fmt.Println("Vous n'avez pas assez d'XP!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		}
	}
	if nom_arbre == 3 {
		if v.xp >= 30 {
			v.attaque += 8
			v.xp -= 30
			fmt.Println("L'attaque a été augmentée de 8!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		} else {
			fmt.Println("Vous n'avez pas assez d'XP!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		}
	}
	if nom_arbre == 4 {
		if v.xp >= 40 {
			v.magie += 5
			v.xp -= 40
			fmt.Println("La magie a été augmentée de 5!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		} else {
			fmt.Println("Vous n'avez pas assez d'XP!")
			time.Sleep(2000 * time.Millisecond)
			os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
			v.Arbre()
		}
	}
	if nom_arbre == 10 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
}
