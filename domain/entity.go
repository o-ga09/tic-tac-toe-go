package domain

type Board struct {
	Board [3][3]string
}

type Koma struct {
	Order int
	X int
	Y int
}

func (b *Board) Init() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b.Board[i][j] = "-"
		}
	}
}