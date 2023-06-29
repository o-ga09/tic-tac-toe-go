package ui

import (
	"fmt"
	"go-tic-tac-toe/domain"
	"io"
	"os"
)

func Input(in io.Reader) string {
	if in == nil {
		in = os.Stdin
	}

	var str string
	fmt.Fscan(in,&str)
	return str
}

func Display(board domain.Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s",board.Board[i][j])
		}
		fmt.Printf("\n")
	}
}