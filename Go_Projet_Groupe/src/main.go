package main

import (
	"fmt"
	"os"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func (p *personnage) loadingmarchand() {
	fmt.Println("LOADING...")
	SlowPrintLoad("[■■■■■■■■■■■■■■■■■■]\n")
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.marchand()
}

func (p *personnage) loadingforgeron() {
	fmt.Println("LOADING...")
	SlowPrintLoad("[■■■■■■■■■■■■■■■■■■]\n")
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	p.forgeron()
}

func SlowPrint(strings ...string) {
	for _, str := range strings {
		for _, char := range str {
			fmt.Print(string(char))
			time.Sleep(40_000000 * time.Nanosecond)
		}
	}
}

func SlowPrintLoad(strings ...string) {
	for _, str := range strings {
		for _, char := range str {
			fmt.Print(string(char))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	var p1 personnage
	var m1 monstre
	i := 0
	p1.Init("Maximator", "Mage noir", 1, 100, 40, []string{"Potion de poison", "Potion de poison", "Potion de soin"}, 100, []string{"Coup de poing"}, []string{}, 0, 8, 10, 19, 16, 150)
	m1.Init("Dragon", 40, 40, 5, 0)
	for i != 100000 {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		p1.menu()
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		m1.fightersquad(&p1)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		i = i + 1
	}
}
