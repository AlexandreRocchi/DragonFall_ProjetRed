package main

import (
	"fmt"
	"time"
)

func (v *personnage) equiper() {
	for i := 0; i <= len(v.inventaire)-1; i++ {
		if v.inventaire[i] == "Chapeau de l'aventurier" {
			v.hpmax = v.hpmax + 10
			v.equipement = append(v.equipement, "Chapeau de l'aventurier")
			v.inventaire[i] = "Emplacement vide"
			v.removeinventory()
			fmt.Println("Chapeau de l'aventurier équipé")
			time.Sleep(1000 * time.Millisecond)
		}
		if v.inventaire[i] == "Tunique de l'aventurier" {
			v.hpmax = v.hpmax + 25
			v.equipement = append(v.equipement, "Tunique de l'aventurier")
			v.inventaire[i] = "Emplacement vide"
			v.removeinventory()
			fmt.Println("Tunique de l'aventurier équipée")
			time.Sleep(1000 * time.Millisecond)
		}
		if v.inventaire[i] == "Bottes de l'aventurier" {
			v.hpmax = v.hpmax + 15
			v.equipement = append(v.equipement, "Bottes de l'aventurier")
			v.inventaire[i] = "Emplacement vide"
			v.removeinventory()
			fmt.Println("Bottes de l'aventurier équipées")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
