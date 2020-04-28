package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/peterh/liner"
)

func main() {
	filePath := flag.String("file", "", "Config file path.")
	flag.Parse()

	configs, err := getConfig(*filePath)
	if err != nil {
		log.Fatalln(err)
	}
	engine := NewEngine(configs)

	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	for {
		if str, err := line.Prompt("> "); err == nil {
			table, err := engine.Translate(str)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(table)
			line.AppendHistory(str)
		} else if err == liner.ErrPromptAborted {
			break
		} else {
			log.Fatalln("Error reading line: ", err)
			break
		}
	}
}
