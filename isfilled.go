package janel

func IsFilled(tab *[3][3]byte) bool {
	for _, t := range tab {
		for _, n := range t {
			if n != 'X' && n != 'O' {
				return false
			}
		}
	}
	return true
}
