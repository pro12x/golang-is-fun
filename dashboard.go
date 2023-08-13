package janel

import "github.com/fatih/color"

func Dashboard(c, min, max, s, l int) {
	color.Cyan("*****************************************************************")
	color.Cyan("   Credit : %d | Min : %d | Max : %d | Score : %d | Level : %d", c, min, max, s, l)
	color.Cyan("*****************************************************************")
}
