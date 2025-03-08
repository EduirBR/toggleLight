package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func findScrollLockPath() (string, error) {
	matches, err := filepath.Glob("/sys/class/leds/input*::scrolllock/brightness")
	if err != nil || len(matches) == 0 {
		return "", fmt.Errorf("could not find scroll lock path")
	}
	return matches[0], nil
}

func main() {
	ledPath, err := findScrollLockPath()
	if err != nil {
		return
	}

	data, err := os.ReadFile(ledPath)
	if err != nil {
		return
	}

	state := strings.TrimSpace(string(data))
	newState := "0"
	if state == "0" {
		newState = "1"
	}

	cmd := exec.Command("bash", "-c", "echo "+newState+" | sudo tee "+ledPath+" > /dev/null")
	cmd.Run()
}
