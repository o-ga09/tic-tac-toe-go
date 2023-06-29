package domain_test

import (
	"go-tic-tac-toe/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInitBoard(t *testing.T) {
	board := &domain.Board{}
	board.Init()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t, "-" ,board.Board[i][j])
		}
	}
}