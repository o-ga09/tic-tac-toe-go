package ui_test

import (
	"bytes"
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/util"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInput(t *testing.T) {
	in := bytes.NewBufferString("hoge")
	driver := &ui.IODriverImpl{}
	got := driver.Input(in)
	assert.Equal(t,"hoge",got)
}

func TestDisplay(t *testing.T) {
	displayFn := func() {
		i := &ui.IODriverImpl{}
		b := [][]string{}
		util.BoardInit(&b)
		driverBoard := ui.DriverBoard{Board: b} 
		i.Display(driverBoard)
	}

	result := util.PickStdout(t, displayFn)
	
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.Equal(t,"---\n---\n---\n",result)
		}
	}
}