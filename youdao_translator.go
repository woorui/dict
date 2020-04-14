package main

import (
	"crypto/sha256"
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

// YoudaoTranslateResult is response of youdao-api
type YoudaoTranslateResult struct {
	ErrorCode   string                     `json:"errorCode"`
	Translation []string                   `json:"translation"`
	Query       string                     `json:"query"`
	Web         []YoudaoTeanslateResultWeb `json:"web"`
	Webdict     string                     `json:"webdict"`
	TransResult []TransResult              `json:"trans_result"`
}

// YoudaoTeanslateResultWeb is sub-struct of YoudaoTranslateResult
type YoudaoTeanslateResultWeb struct {
	Phonetic string   `json:"phonetic"`
	Explains []string `json:"explains"`
}

// baseurl	https://openapi.youdao.com/api
func youydaoTranslator(text string) {

}

// timestamp := strconv.Itoa(time.Now().Second())
// Passing the timestamp make function testable
func genQuery(text string, timestamp string) {
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

	u, err := url.Parse(baseurl)
	if err != nil {
		return
	}
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.Path = path
	u.RawQuery = q.Encode()
	return
}

func genInput(p string) string {
	b := []byte(p)
	if len(b) >= 20 {
		return string(b[:10]) + strconv.Itoa(len(b)) + string(b[len(b)-10:])
	}
	return string(b)
}

func genSign(str string) string {
	hash := sha256.Sum256([]byte(str))
	return string(hash[:])
}
