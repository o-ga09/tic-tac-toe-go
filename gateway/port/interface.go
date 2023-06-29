package port

import (
	"go-tic-tac-toe/domain"
	"io"
)

type GameInterface interface {
	IsWin() bool
	Display(b domain.Board)
	Input(io.Reader) domain.Koma
}