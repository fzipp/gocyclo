package main_test

import (
	"os/exec"
	"testing"
)

func TestHelp(t *testing.T) {
	exec.Command("gocycle", "-h").Run()
}
