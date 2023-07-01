package ui

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type IODriver interface {
	Input(io.Reader) (DriverKoma, error)
	Display(DriverBoard) error
}

type IODriverImpl struct {}

func PrioviderUIDriver() IODriver {
	return &IODriverImpl{}
}

func (i *IODriverImpl) Input(in io.Reader) (DriverKoma, error) {
	if in == nil {
		in = os.Stdin
	}

	var str string
	fmt.Fscan(in,&str)

	args := strings.Split(str, ",")
	if len(args) != 2 {
		return DriverKoma{}, errors.New("座標入力値フォーマットが不正")
	}

	x, errX := strconv.Atoi(args[0])
	y, errY := strconv.Atoi(args[1])
	if errX != nil || errY != nil {
		return DriverKoma{}, errors.New("座標入力値の型が違います want int but string")
	}
	return DriverKoma{X: x,Y: y}, nil
}

func (i *IODriverImpl) Display(board DriverBoard) error {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s",board.Board[i][j])
		}
		fmt.Printf("\n")
	}
	return nil
}

type DriverBoard struct {
	Board [][]string
}

type DriverKoma struct {
	X int
	Y int
}