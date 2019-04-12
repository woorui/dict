package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func askQuestion(reader *bufio.Reader, question string, allowEmpty bool) string {
	fmt.Println(question)
	input, err := getTrimInput(reader, allowEmpty)
	if err != nil {
		fmt.Println(err.Error())
		return askQuestion(reader, question, allowEmpty)
	}
	return input
}

func getTrimInput(reader *bufio.Reader, allowEmpty bool) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil || (str == "\n" && allowEmpty == false) {
		if err != nil {
			return "", err
		}
		return "", errors.New("The input is not allowed empty")
	}

	return strings.Trim(str, " \n"), nil
}

func askInformation(reader *bufio.Reader) *Information {
	askQuestion(reader, "You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index", true)
	appid := askQuestion(reader, "Input your baidu appid", false)
	secret := askQuestion(reader, "Input your baidu secret", false)
	return &Information{Appid: appid, Secret: secret}
}

func formatAndPrintRes(tr []TransResult) {
	fmt.Println(formatPes(tr))
}

func formatPes(tr []TransResult) string {
	var s string
	if len(tr) == 0 {
		return ""
	}
	for _, v := range tr {
		s += fmt.Sprintf("|> %s: %s", v.Src, v.Dst)
	}
	return s
}
