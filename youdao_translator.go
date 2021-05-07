package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type youdaoTranslator struct {
	name      string
	client    *http.Client
	baseurl   string
	appKey    string
	appSecret string
}

func newYoudaoTranslator(client *http.Client, baseurl, str string) *youdaoTranslator {
	arr := strings.Split(str, "-")
	if len(arr) != 2 {
		panic("config string error")
	}
	appKey := arr[0]
	appSecret := arr[1]
	return &youdaoTranslator{
		name:      youdaoName,
		client:    client,
		baseurl:   baseurl,
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (translator *youdaoTranslator) GetName() string {
	return translator.name
}

func (translator *youdaoTranslator) Translate(text string) ([]Translation, error) {
	timestamp := strconv.Itoa(int(time.Now().Unix()))
	url := translator.genRequestURL(text, timestamp, uuid.New().String())

	ytr, err := translator.doRequest(url, text)
	if err != nil {
		return nil, err
	}
	return youdaoTransformer(ytr), nil
}

// genRequestURL generate url request by client
// Passing the timestamp make function testable
func (translator *youdaoTranslator) genRequestURL(text, timestamp, salt string) string {
	query := map[string]string{
		"q":        text,
		"appKey":   translator.appKey,
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
	input := translator.appKey + genInput(text) + salt + timestamp + translator.appSecret
	hash := sha256.New()
	hash.Write([]byte(input))
	query["sign"] = fmt.Sprintf("%x", hash.Sum(nil)) // 16 进制

	u, _ := url.Parse(youdaoURL)
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (translator *youdaoTranslator) doRequest(url string, text string) (YoudaoTranslateResult, error) {
	var t YoudaoTranslateResult
	if text == "" {
		return t, nil
	}
	client := translator.client
	body, err := HTTPGetRequest(client, url)
	if err != nil {
		return t, nil
	}
	return t, unmarshalYoudaoResBody(&t, body)
}

func unmarshalYoudaoResBody(t *YoudaoTranslateResult, body []byte) error {
	if err := json.Unmarshal(body, &t); err != nil {
		return err
	}
	if t.ErrorCode != "" {
		return youdaoErrCodeMessage[t.ErrorCode]
	}
	return nil
}

func genInput(p string) string {
	b := []rune(p)
	if len(b) > 20 {
		return string(b[:10]) + strconv.Itoa(len(b)) + string(b[len(b)-10:])
	}
	return string(b)
}

func youdaoTransformer(t YoudaoTranslateResult) []Translation {
	var arr []Translation
	item := Translation{
		DataSource: youdaoName,
		Src:        t.Query,
		Dst:        strings.Join(t.Translation, ", "),
		Phonetic:   t.Basic.Phonetic,
		Explain:    strings.Join(t.Basic.Explains, ", "),
	}
	arr = append(arr, item)
	return arr
}
