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
	rep   string // Pour récupérer la réponse de l'utilisateur
	num   int    // Le nombre mystère
	nb    int    // Pour récupérer le nombre de l'utilisateur
	check bool   // Check user's answer
)

func main() {
	var player Player
	min, max := 0, 100

	fmt.Println("Type your pseudo : ")
	fmt.Scan(&player.name)

	for true {
		num = rand.Intn(max-min) + min
		fmt.Println("Guess the number between", min, "and", max)
		for true {
			//fmt.Println("Guess the number between", min, "and", max)
			fmt.Scan(&nb)
			if nb < min || nb > max {
				CleanScreen()
				fmt.Println("Valid number is between", min, "and", max)
				continue
			}
			if nb > num {
				CleanScreen()
				fmt.Println("The mystery number is less than", nb)
			} else if nb < num {
				CleanScreen()
				fmt.Println("The mystery number is greater than", nb)
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
			//fmt.Println("Guess the number between", min, "and", max)
		}
		break
	}
	CleanScreen()
	GameOver(player.name, player.level, player.score, player.games)
}

func CleanScreen() {
	screen.Clear()
	screen.MoveTopLeft()
}

func GameOver(n string, l, s, g int) {
	fmt.Println("*****************************************")
	fmt.Println("*\tPlayer :", n, "\t\t\t*")
	fmt.Println("*\tLevel :", l, "\t\t\t*")
	fmt.Println("*\tScore :", s, "\t\t\t*")
	fmt.Println("*\tYou played :", g, "times\t\t*")
	fmt.Println("*****************************************")
}
