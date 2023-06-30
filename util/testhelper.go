package util

import (
	"bytes"
	"os"
	"testing"
)

func PickStdout(t *testing.T, fnc func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	fnc()
	w.Close()
	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v", err)
	}
	s := buffer.String()
	return s
}

func BoardInit(b *[][]string) {
	*b = make([][]string, 3)
	for i := 0; i < 3; i++ {
		(*b)[i] = make([]string, 3)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			(*b)[i][j] = "-"
		}
	}
}