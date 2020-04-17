package main

import (
	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	for {
		if input, err := line.Prompt("> "); err == nil {
			// TODO do request
		}
		break
	}
}
