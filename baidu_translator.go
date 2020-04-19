package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
)

type baiduTranslator struct {
	name    string
	client  *http.Client
	baseurl string
	appID   string
	secret  string
}

func newBaiduTranslator(client *http.Client, baseurl, appID, secret string) baiduTranslator {
	return baiduTranslator{
		name:    "百度",
		client:  client,
		baseurl: baseurl,
		appID:   appID,
		secret:  secret,
	}
}

// salt := strconv.Itoa(rand.Int() * 1000)
func (translator *baiduTranslator) genRequestURL(text string, salt string) (string, error) {
	sign := generateHashSign(translator.appID, text, salt, translator.secret)
	query := map[string]string{
		"q":     text,
		"appid": translator.appID,
		"salt":  salt,
		"sign":  sign,
		"from":  "auto",
	}

	if textContainChinese(text) {
		query["to"] = "en"
	} else {
		query["to"] = "zh"
	}

	u, err := url.Parse(translator.baseurl)
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
	t := BaiduTranslateResult{}
	client := translator.client
	salt := strconv.Itoa(rand.Int() * 1000)
	url, err := translator.genRequestURL(text, salt)
	if err != nil {
		return t, nil
	}
	body, err := HTTPGetRequest(client, url)
	if err != nil {
		return t, nil
	}
	return unmarshalBaiduResBody(t, body)
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

func unmarshalBaiduResBody(t BaiduTranslateResult, body []byte) (BaiduTranslateResult, error) {
	if err := json.Unmarshal(body, &t); err != nil {
		return t, err
	}
	if t.ErrorCode != "" {
		return t, baiduErrCodeMessage[t.ErrorCode]
	}
	return t, nil
}
