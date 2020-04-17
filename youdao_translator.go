package main

import (
	"crypto/sha256"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// YoudaoTranslateResult is response of youdao-api
type YoudaoTranslateResult struct {
	ErrorCode   string   `json:"errorCode"`
	Translation []string `json:"translation"`
	Query       string   `json:"query"`
	Basic       struct {
		Phonetic string   `json:"phonetic"`
		Explains []string `json:"explains"`
	} `json:"basic"`
	Web         []YoudaoTranslateResultWeb `json:"web"`
	Webdict     string                     `json:"webdict"`
	TransResult []TransResult              `json:"trans_result"`
}

// YoudaoTranslateResultWeb is sub-struct of YoudaoTranslateResult
type YoudaoTranslateResultWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}

func youydaoTranslator(text string) {

}

// timestamp := strconv.Itoa(time.Now().Second())
// Passing the timestamp make function testable
func genRequestURLs(text string, timestamp string) (string, error) {
	salt := uuid.New().String()
	query := map[string]string{
		"q":        text,
		"appKey":   "appKey",
		"salt":     salt,
		"signType": "v3",
		"curtime":  timestamp,
	}
	if wordsContainChinese(text) {
		query["from"] = "zh-CHS"
		query["to"] = "en"
	} else {
		query["to"] = "zh-CHS"
		query["from"] = "en"
	}
	input := "appKey" + genInput(text) + salt + timestamp + "appSecret"
	hash := sha256.Sum256([]byte(input))
	query["sign"] = string(hash[:])

	u, err := url.Parse(youdaoURL)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func doRequests(text string) (transRes, error) {
	var raw transRes
	if text == "" {
		return raw, nil
	}
	client := &http.Client{}

	url, err := genRequestURLs(text, strconv.Itoa(time.Now().Second()))
	if err != nil {
		return raw, nil
	}

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return raw, err
	}

	resp, err := client.Do(r)
	if err != nil {
		return raw, err
	}

	return parseResponse(resp.Body)
}

func genInput(p string) string {
	b := []byte(p)
	if len(b) >= 20 {
		return string(b[:10]) + strconv.Itoa(len(b)) + string(b[len(b)-10:])
	}
	return string(b)
}
