package janel

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type Player struct {
	name  string
	score int
	level int
	games int
}

var (
	rep    string
	num    int
	nb     int
	check  bool
	rule   int
	credit int
)

func Mystery() {
	CleanScreen()
	Rules()
	fmt.Scan(&rule)
	if rule == 1 {
		var player Player
		min, max := 0, 100
		player.level = 1
		player.score = 0
		credit = 7

		CleanScreen()
		Dashboard(credit, min, max, player.score, player.level)
		color.Cyan("Type your pseudo : ")
		fmt.Scan(&player.name)

		for true {
			num = rand.Intn(max-min) + min
			color.Cyan("Guess the number between %d and %d", min, max)
			for true {
				if credit == 0 {
					break
				}
				fmt.Scan(&nb)
				if nb < min || nb > max {
					CleanScreen()
					Dashboard(credit, min, max, player.score, player.level)
					color.Red("Valid number is between %d and %d", min, max)
					continue
				}
				if nb > num {
					credit -= 1
					CleanScreen()
					Dashboard(credit, min, max, player.score, player.level)
					color.Cyan("The mystery number is less than %d", nb)
				} else if nb < num {
					credit -= 1
					CleanScreen()
					Dashboard(credit, min, max, player.score, player.level)
					color.Cyan("The mystery number is greater than %d", nb)
				} else {
					credit -= 1
					CleanScreen()
					Dashboard(credit, min, max, player.score, player.level)
					color.Green("Congrat! You found %d after %d tries", num, 5-credit)
					player.score = player.score + (player.level * 5)
					for true {
						color.Cyan("Do you want to continue? (Y/n)")
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
						credit = 7
						CleanScreen()
						Dashboard(credit, min, max, player.score, player.level)
					} else {
						break
					}
				}
				player.games++
				// Dashboard(credit, min, max, player.score, player.level)
				//fmt.Println("Guess the number between", min, "and", max)
			}
			break
		}
		CleanScreen()
		GameOver(player.name, player.level, player.score, player.games)
		time.Sleep(5 * time.Second)
		// CleanScreen()
		if player.level < 5 {
			Looser()
		} else {
			Winner()
		}
	} else {
		return
	}
}
