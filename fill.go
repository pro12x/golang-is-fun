package janel

func Fill(nb int, val byte, tab1 *[3][3]byte, tab2 [3][3]int) (int, int, bool) {
	for i, t := range tab2 {
		for j, n := range t {
			if nb == n && (tab1[i][j] != 'X' && tab1[i][j] != 'O') {
				tab1[i][j] = val
				return i, j, true
			}
		}
	}
	return 0, 0, false
}
