package main

import "fmt"

func (v *personnage) takePot() {
	for i := 0; i <= len(v.inventaire)-1; i++ {
		if v.inventaire[i] == "Potion de soin" {
			v.hpactuels += 50
			v.inventaire[i] = "Emplacement vide"
			v.removeinventory()
			break
		}
	}
	if v.hpactuels > v.hpmax {
		v.hpactuels = v.hpmax
	}
	fmt.Println("vie = ", v.hpactuels, "/ 100 ")
}
