package ui_test

import (
	"bufio"
	"bytes"
	ui "go-tic-tac-toe/UI"
	"go-tic-tac-toe/util"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInput(t *testing.T) {
	in := bytes.NewBufferString("1,2")

	// スペースで区切られた文字列を読み込む
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	driver := ui.PrioviderUIDriver()
	got,err := driver.Input(in)
	assert.Equal(t,ui.DriverKoma{X: 1,Y: 2},got)
	assert.Equal(t,nil,err)
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