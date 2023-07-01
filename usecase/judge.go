package usecase

import (
	"errors"
	"go-tic-tac-toe/domain"
	"go-tic-tac-toe/gateway/port"
)

type GameService struct {
	gameInterface port.GameInterface
	board *domain.Board
}

const (
	ROW_NUM = 3
	COLUMN_NUM= 3
	EMPTY = "-"
)

func ProviderGameDriver(gameGateway port.GameInterface) GameService {
	return GameService{
		gameInterface: gameGateway,
	}
}

func (g *GameService) Init() *domain.Board {
	g.board = &domain.Board{Board: [3][3]string{}}
	g.board.Init()
	return g.board
}

func (g *GameService) IsWin(board domain.Board) bool {
	if checkVertical(&board) || checkHorizon(&board) || checkCross(&board) {
		return true
	}

	return false
}

func (g *GameService) Display(board *domain.Board) error {
	err := g.gameInterface.Display(board)

	return err
}

func (g *GameService) Input(player int) (domain.Koma, error) {
	res, err := g.gameInterface.Input(player)
	if err != nil {
		return domain.Koma{}, err
	}
	if res.X >= ROW_NUM || res.Y >= COLUMN_NUM {
		return domain.Koma{}, err
	}
	if !isEmpty(g.board,res) {
		return domain.Koma{},errors.New("koma is exist")
	} 

	switch (player) {
	case 1:
		g.board.Board[res.X][res.Y] = "○"
	case 2:
		g.board.Board[res.X][res.Y] = "×"
	default :
	}
	return domain.Koma{Order: res.Order,X: res.X,Y: res.Y}, nil
}

func checkVertical(board *domain.Board) bool {
	for i := 0;i < COLUMN_NUM;i++ {
		if ((((*board).Board[0][i]) == (*board).Board[1][i] && (*board).Board[1][i] == (*board).Board[2][i] && (*board).Board[0][i] == (*board).Board[2][i]) && (*board).Board[0][i] != EMPTY){
			return true
		}
	}
	return false
}

func checkHorizon(board *domain.Board) bool {
	for i := 0;i < ROW_NUM;i++ {
		if (((*board).Board[i][0] == (*board).Board[i][1] && (*board).Board[i][1] == (*board).Board[i][2] && (*board).Board[i][0] == (*board).Board[i][2]) && (*board).Board[i][0] != EMPTY){
			return true
		}
	}
	return false
}

func checkCross(board *domain.Board) bool {
	if (((*board).Board[0][0] == (*board).Board[1][1] && (*board).Board[1][1] == (*board).Board[2][2] && (*board).Board[0][0] == (*board).Board[2][2]) && (*board).Board[0][0] != EMPTY){
		return true
	}else if (((*board).Board[0][2] == (*board).Board[1][1] && (*board).Board[1][1] == (*board).Board[2][0] && (*board).Board[0][2] == (*board).Board[2][0]) && (*board).Board[0][2] != EMPTY){
		return true
	}
	return false
}

func isEmpty(board *domain.Board, pos domain.Koma) bool {
	return (*board).Board[pos.X][pos.Y] == EMPTY
}