package janel

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

/* ==================== Global Variables ===================== */
var (
	board = [3][3]byte{
		{'1', '2', '3'},
		{'4', '5', '6'},
		{'7', '8', '9'},
	}

	hint = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	val    int
	game   bool = true
	player byte
	c      int = 1
	f      bool
)

func Morpion() {
	for !IsFilled(&board) && !Check(&board) {
		Display(board)
		color.Cyan("Player %d :", c)
		fmt.Scan(&val)
		if val < 1 || val > 9 {
			CleanScreen()
			color.Red("Only digits between 1 and 9")
		} else {
			CleanScreen()
			if game {
				player = 'X'
				game = false
			} else {
				player = 'O'
				game = true
			}
			Fill(val, player, &board, hint)
			c = Players(player)
		}
	}

	CleanScreen()
	Display(board)
	time.Sleep(2 * time.Second)
	if IsFilled(&board) {
		if Check(&board) {
			color.Green("Winner is : Player_%c", player)
		} else {
			color.Yellow("Draw")
		}
	} else {
		color.Green("Winner is : Player_%c", player)
	}
}
