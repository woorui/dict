package main

import (
	"fmt"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	for {
		if text, err := line.Prompt("> "); err == nil {
			// TODO do request
			fmt.Println(text)
		}
		break
	}
}
