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