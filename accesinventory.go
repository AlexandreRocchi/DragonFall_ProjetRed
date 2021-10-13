package main

import "fmt"

func (v *personnage) accesInventory() {
	if len(v.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	} else {
		fmt.Println("Objets :", v.inventaire)
		fmt.Println("Sorts :", v.sorts)
	}
}
