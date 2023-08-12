package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Player struct {
	name  string
	score int
	level int
}

func main() {
	//My variables
	var player Player
	var in string
	var test int
	//var number int
	var score int
	var level int
	score = 120
	level = 12

	//Assigning
	player.name = "Janel"
	player.score = score
	player.level = level

	for test <= 0 || test > 100 {
		fmt.Scan(&in)
		test, _ = strconv.Atoi(in)
		if test == 0 || test < 0 || test > 100 {
			fmt.Println("Type a valid number please !")
		}
	}

	//Display
	/* fmt.Println(player.name)
	fmt.Println(player.score)
	fmt.Println(player.level)
	fmt.Println(number) */
	fmt.Println(test)
	fmt.Println(rand.Intn(100-1) + 1)
}
