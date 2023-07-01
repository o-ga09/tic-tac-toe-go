package port

import (
	"go-tic-tac-toe/domain"
)

type GameInterface interface {
	IsWin(domain.Board) bool
	Display(b *domain.Board) error
	Input(int) (domain.Koma, error)
}