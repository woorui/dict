package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input words to translate")
	for scanner.Scan() {
		input := scanner.Text()
		tr, err := doRequest(appid, secret, input)
		if err != nil {
			if tr.ErrorCode == "52003" || tr.ErrorCode == "54001" {
				removeInformation()
			}
			log.Fatalln(err)
		}
		formatAndPrintRes(tr.TransResult)
	}
}
