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
	client http.Client
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

func (translator *baiduTranslator) doRequest(text string) (BaiduTranslateResult, error) {
	var t BaiduTranslateResult
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

	return unmarshalBaiduResBody(t, resp.Body)
}

// BaiduTranslateResult is response body from remote api
type BaiduTranslateResult struct {
	ErrorCode   string        `json:"error_code"`
	ErrorMsg    string        `json:"error_msg"`
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}

// TransResult is type of BaiduTranslateResult.TransResult
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

func wordsContainChinese(text string) bool {
	return utf8.RuneCountInString(text) != len(text)
}

func unmarshalBaiduResBody(t BaiduTranslateResult, rc io.ReadCloser) (BaiduTranslateResult, error) {
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return t, err
	}
	if err := json.Unmarshal(body, &t); err != nil {
		return t, err
	}
	if t.ErrorCode != "" {
		return t, errors.New(errCodeMessage[t.ErrorCode])
	}
	return t, nil
}
