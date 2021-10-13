package main

import (
	"fmt"
	"os"
	"time"
)

func (v *monstre) fightersquad(p *personnage) {
	v.tour_monstre = v.tour_monstre + 1
	if v.hpactuels_monstre > 0 {
		fmt.Println("Tour ", v.tour_monstre)
	}
	if v.hpactuels_monstre <= 0 {
		fmt.Print("Vous avez gagné le dragon en ", v.tour_monstre, "tour(s)")
		v.hpmax_monstre = v.hpmax_monstre * 2
		v.hpactuels_monstre = v.hpmax_monstre
		v.attack_monstre = v.attack_monstre * 2
		p.argent += 30
		p.xp += 65
		time.Sleep(2000 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p.menu()
	} else {
		v.charturn(p)
	}
}
