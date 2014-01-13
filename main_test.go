package main

import (
	"github.com/yosssi/gocmd"
	"os/exec"
	"testing"
)

const (
	targetFileName = "main_test.go"
)

func TestConvOutputToStrings(t *testing.T) {
	output, _ := gocmd.Pipe(exec.Command("ls"), exec.Command("grep", targetFileName))
	s := convOutputToStrings(output)
	if len(s) != 1 || s[0] != targetFileName {
		t.Error("The return value should be [" + targetFileName + "].")
	}

	output, _ = gocmd.Pipe(exec.Command("ls"), exec.Command("grep", "abcdefghijklmnopqrstuvwxyz"))
	s = convOutputToStrings(output)
	if len(s) != 0 {
		t.Error("The return value should be [].")
	}
}
