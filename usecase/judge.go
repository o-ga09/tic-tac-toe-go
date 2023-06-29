package usecase

import (
	"fmt"
	"go-tic-tac-toe/domain"
	"go-tic-tac-toe/gateway/port"
	"io"
	"os"
	"strconv"
	"strings"
)

type GameService struct {
	gameInterface port.GameInterface
}

func ProviderGameDriver(gameGateway port.GameInterface) GameService {
	return GameService{gameInterface: gameGateway}
}

func (g *GameService) IsWin() {

}

func (g *GameService) Display(board domain.Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%s",board.Board[i][j])
		}
		fmt.Printf("\n")
	}
}

func (g *GameService) Input(in io.Reader) domain.Koma {
	if in == nil {
		in = os.Stdin
	}

	var str string
	fmt.Fscan(in,&str)
	pos := strings.Split(str, " ")
	x, _ := strconv.Atoi(pos[0])
	y, _ := strconv.Atoi(pos[1])
	return domain.Koma{
		X: x,
		Y: y,	
	}
}