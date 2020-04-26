package main

import (
	"fmt"
	"log"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)
	for {
		if str, err := line.Prompt("> "); err == nil {
			fmt.Print(str)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Fatalln("Error reading line: ", err)
		}
	}
}
