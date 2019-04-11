package main

import (
	"bufio"
	"fmt"
	"strings"
)

func getTrimInput(reader *bufio.Reader, question string, allowEmpty bool) string {
	fmt.Println(question)
	str, err := reader.ReadString('\n')
	if err != nil || (str == "\n" && allowEmpty == false) {
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("The input is not allowed empty")
		return getTrimInput(reader, question, false)
	}

	return strings.Trim(str, " \n")
}

func askInformation(reader *bufio.Reader) *Information {
	getTrimInput(reader, "You need baidu secret and appid, apply link: http://api.fanyi.baidu.com/api/trans/product/index", true)
	appid := getTrimInput(reader, "Input your baidu appid", false)
	secret := getTrimInput(reader, "Input your baidu secret", false)
	return &Information{Appid: appid, Secret: secret}
}

func formatAndPrintRes(tr []TransResult) {
	var s string
	if len(tr) == 0 {
		return
	}
	for _, v := range tr {
		s += fmt.Sprintf("|> %s: %s", v.Src, v.Dst)
	}
	fmt.Println(s)
}
