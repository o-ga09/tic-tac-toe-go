package ui_test

import (
	"bytes"
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/domain"
	"go-tic-tac-toe/util"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInput(t *testing.T) {
	in := bytes.NewBufferString("hoge")
	got := ui.Input(in)
	assert.Equal(t,"hoge",got)
}

func TestDisplay(t *testing.T) {
	board := &domain.Board{}
	board.Init()
	got := util.PickStdout(t, ui.Display)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t,"---\n---\n---\n",got)
		}
	}
}