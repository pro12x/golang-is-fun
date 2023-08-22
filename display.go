package janel

import "fmt"

func Display(tab [3][3]byte) {
	for _, t := range tab {
		fmt.Println("-------------")
		fmt.Print("|")
		for _, n := range t {
			fmt.Printf(" %c |", n)
		}
		fmt.Println()
	}
	fmt.Println("-------------")
}
