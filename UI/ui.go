package ui

import (
	"fmt"
	"io"
	"os"
)

type IODriver interface {
	Input(io.Reader)
	Display(DriverBoard)
}

type IODriverImpl struct {}

func (i *IODriverImpl) Input(in io.Reader) string {
	if in == nil {
		in = os.Stdin
	}

	var str string
	fmt.Fscan(in,&str)
	return str
}

func (i *IODriverImpl) Display(board DriverBoard) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s",board.Board[i][j])
		}
		fmt.Printf("\n")
	}
}

type DriverBoard struct {
	Board [][]string
}