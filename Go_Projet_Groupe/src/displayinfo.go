package main

import "fmt"

func (v personnage) DisplayInfo() {
	fmt.Println("---------")
	fmt.Println("Son nom est:", v.nom)
	fmt.Println("Sa classe est:", v.classe)
	fmt.Println("Il est de niveau est:", v.niveau)
	fmt.Println("Sa vie max est:", v.hpmax)
	fmt.Println("Sa vie actuelle est:", v.hpactuels)
	fmt.Println("Inventaire:", v.inventaire)
	fmt.Println("Pièces:", v.argent)
}
