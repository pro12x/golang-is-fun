package janel

import "github.com/fatih/color"

func GameOver(n string, l, s, g int) {
	color.Cyan("*****************************************************************")
	color.Cyan("\tPlayer : %s", n)
	color.Cyan("\tLevel : %d", l)
	color.Cyan("\tScore : %d", s)
	color.Cyan("\tYou played : %d times", g)
	color.Cyan("*****************************************************************")
}
