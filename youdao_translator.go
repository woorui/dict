package main

import (
	"crypto/sha256"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type youdaoTranslator struct {
	name    string
	client  *http.Client
	baseurl string
	appID   string
	secret  string
}

// genRequestURL generate url request by client
// timestamp := strconv.Itoa(time.Now().Second())
// Passing the timestamp make function testable
func (translator *youdaoTranslator) genRequestURL(text string, timestamp string) (string, error) {
	salt := uuid.New().String()
	query := map[string]string{
		"q":        text,
		"appKey":   "appKey",
		"salt":     salt,
		"signType": "v3",
		"curtime":  timestamp,
	}
	if textContainChinese(text) {
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

func (translator *youdaoTranslator) doRequest(text string) (YoudaoTranslateResult, error) {
	var t YoudaoTranslateResult
	if text == "" {
		return t, nil
	}
	client := translator.client
	timestamp := strconv.Itoa(time.Now().Second())
	url, err := translator.genRequestURL(text, timestamp)
	if err != nil {
		return t, nil
	}
	body, err := HTTPGetRequest(client, url)
	if err != nil {
		return t, nil
	}
	return unmarshalYoudaoResBody(t, body)
}

func unmarshalYoudaoResBody(t YoudaoTranslateResult, body []byte) (YoudaoTranslateResult, error) {
	if err := json.Unmarshal(body, &t); err != nil {
		return t, err
	}
	if t.ErrorCode != "" {
		return t, baiduErrCodeMessage[t.ErrorCode]
	}
	return t, nil
}

func genInput(p string) string {
	b := []byte(p)
	if len(b) >= 20 {
		return string(b[:10]) + strconv.Itoa(len(b)) + string(b[len(b)-10:])
	}
	return string(b)
}
