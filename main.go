package main

import (
	"bufio"
	"log"
	"os"

	"github.com/peterh/liner"
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	reader := bufio.NewReader(os.Stdin)
	information, err := getInformation()
	if err != nil {
		information = askInformation(reader)
		if err := saveInformation(information); err != nil {
			log.Fatalln("save appid and secret error, exit!")
		}
	}

	appid := information.Appid
	secret := information.Secret

	for {
		if input, err := line.Prompt("> "); err == nil {
			tr, err := doRequest(appid, secret, input)
			if err != nil {
				if tr.ErrorCode == "52003" || tr.ErrorCode == "54001" {
					removeInformation()
				}
				log.Fatalln(err)
			}
			formatAndPrintRes(tr.TransResult)
			line.AppendHistory(input)
		} else if err == liner.ErrPromptAborted {
			log.Print("Aborted")
			break
		} else {
			log.Print("Error reading line: ", err)
			break
		}
	}
}
