package main

import (
	"fmt"
	"math/rand"

	"github.com/inancgumus/screen"
)

type Player struct {
	name  string
	score int
	level int
	games int
}

var (
	rep   string // Stock user answer
	num   int    // Stock mystery number
	nb    int    // user's number
	check bool   // Check user's answer
)

func main() {
	var player Player
	min, max := 0, 100

	fmt.Println("Type your pseudo : ")
	fmt.Scan(&player.name)

	for true {
		num = rand.Intn(max-min) + min
		for true {
			fmt.Println("Guess the number between", min, "and", max)
			fmt.Scan(&nb)
			if nb > num {
				CleanScreen()
				fmt.Println("It's smaller")
			} else if nb < num {
				CleanScreen()
				fmt.Println("It's bigger")
			} else {
				CleanScreen()
				fmt.Println("Congrat! You found", num)
				for true {
					fmt.Println("Do you want to continue? (Y/n)")
					fmt.Scan(&rep)
					CleanScreen()
					if rep == "N" || rep == "n" {
						check = false
						break
					} else if rep == "Y" || rep == "y" {
						check = true
						break
					}
				}
				if check {
					player.level++
					min += 50
					max += 100
					num = rand.Intn(max-min) + min
				} else {
					break
				}
			}
		}
		break
	}
	fmt.Println("Thank you", player.name, "!")
	fmt.Println("Level :", player.level)
	fmt.Println("Score", player.score)
	fmt.Println("Thank you", player.name)
	fmt.Println("Thank you", player.name)
}

func CleanScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}
