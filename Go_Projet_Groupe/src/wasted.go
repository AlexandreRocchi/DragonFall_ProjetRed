package main

import (
	"fmt"
	"os"
	"time"
)

func (v *personnage) Wasted() {
	if v.hpactuels <= 0 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println("Vous êtes mort            ")
		fmt.Println("                          ")
		fmt.Println("   _________              ")
		fmt.Println("  /         \\            ")
		fmt.Println(" /    RIP    \\           ")
		fmt.Println(" |           |            ")
		fmt.Println(" |           |            ")
		fmt.Println(" |  (.---.)  |            ")
		fmt.Println(" |__,|X X|,__|            ")
		time.Sleep(3 * time.Second)
		v.hpactuels = v.hpmax / 2
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		v.menu()
	}
}
