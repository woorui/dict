package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"unicode/utf8"
)

type baiduTranslator struct {
	url    string
	appID  string
	secret string
}

func (translator *baiduTranslator) genRequestURL(text string) (string, error) {
	salt := strconv.Itoa(rand.Int() * 1000)
	sign := generateHashSign(translator.appID, text, salt, translator.secret)
	query := map[string]string{
		"q":     text,
		"appid": translator.appID,
		"salt":  salt,
		"sign":  sign,
		"from":  "auto",
	}

	if wordsContainChinese(text) {
		query["to"] = "en"
	} else {
		query["to"] = "zh"
	}

	u, err := url.Parse(translator.url)
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

func (translator *baiduTranslator) doRequest(text string) (translation, error) {
	var t translation
	client := &http.Client{}
	url, err := translator.genRequestURL(text)
	if err != nil {
		return t, nil
	}
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return t, err
	}
	resp, err := client.Do(r)
	if err != nil {
		return t, err
	}

	return t, nil
}

// translation is response body from remote api
type translation struct {
	ErrorCode   string        `json:"error_code"`
	ErrorMsg    string        `json:"error_msg"`
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}

// TransResult is translation TransResult field
type TransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func generateHashSign(appid, q, salt, secret string) string {
	var buffer bytes.Buffer
	for _, v := range [4]string{appid, q, salt, secret} {
		buffer.WriteString(v)
	}
	concatStr := buffer.String()
	hasher := md5.New()
	hasher.Write([]byte(concatStr))
	sign := hex.EncodeToString(hasher.Sum(nil))

	return sign
}

func doRequest(appid, secret, word string) (transRes, error) {
	var raw transRes
	if word == "" {
		return raw, nil
	}
	client := &http.Client{}
	salt := strconv.Itoa(rand.Int() * 1000)
	sign := generateHashSign(appid, word, salt, secret)

	url, err := genRequestURL(baseurl, path, word, appid, salt, sign)
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

func wordsContainChinese(text string) bool {
	return utf8.RuneCountInString(text) != len(text)
}

// genRequestURL generator URL to request api
func genRequestURL(baseurl string, path string, word, appid, salt, sign string) (string, error) {
	query := map[string]string{
		"q":     word,
		"appid": appid,
		"salt":  salt,
		"sign":  sign,
		"from":  "auto",
	}

	if wordsContainChinese(word) {
		query["to"] = "en"
	} else {
		query["to"] = "zh"
	}

	u, err := url.Parse(baseurl)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range query {
		q.Set(k, v)
	}
	u.Path = path
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// parseResponse parse the response of translator
func parseResponse(rc io.ReadCloser) (translation, error) {
	var raw translation

	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return raw, err
	}

	if err := json.Unmarshal(body, &raw); err != nil {
		return raw, err
	}

	if raw.ErrorCode != "" {
		return raw, errors.New(errCodeMessage[raw.ErrorCode])
	}

	return raw, nil
}
