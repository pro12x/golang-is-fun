package janel

func Check(tab *[3][3]byte) bool {
	if (tab[0][0] == tab[0][1] && tab[0][0] == tab[0][2]) || (tab[1][0] == tab[1][1] && tab[1][0] == tab[1][2]) || (tab[2][0] == tab[2][1] && tab[2][0] == tab[2][2]) || (tab[0][0] == tab[1][0] && tab[0][0] == tab[2][0]) || (tab[0][1] == tab[1][1] && tab[0][1] == tab[2][1]) || (tab[0][2] == tab[1][2] && tab[0][2] == tab[2][2]) || (tab[0][0] == tab[1][1] && tab[0][0] == tab[2][2]) || (tab[0][2] == tab[1][1] && tab[0][2] == tab[2][0]) {
		return true
	}
	return false
}
