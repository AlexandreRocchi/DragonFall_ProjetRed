package main

import (
	"fmt"
	"os"
	"time"
)

func (v *personnage) spellbook(spell string, price int) {
	if v.argent > price {
		v.sorts = append(v.sorts, spell)
	} else {
		fmt.Println("Vous n'avez pas assez pas d'argent")
		time.Sleep(2 * time.Second)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
}
