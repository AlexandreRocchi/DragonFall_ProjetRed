package main

import (
	"fmt"
	"time"
)

func (v *personnage) poisonPot() {
	n := 1
	for a := 0; a <= len(v.inventaire)-1; a++ {
		if v.inventaire[a] == "Potion de poison" {
			n = n + 1
			if n == 2 {
				v.inventaire[a] = "Emplacement vide"
				v.removeinventory()
				for t := 0; t < 3; t++ {
					v.hpactuels -= 10
					fmt.Println("*ouch* -10 points de vie")
					time.Sleep(1 * time.Second)
					n = n + 2
				}
			}
		}
	}
	fmt.Println("vie = ", v.hpactuels, "/ 100 ")
}
