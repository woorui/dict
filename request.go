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

const baseurl = "http://api.fanyi.baidu.com"
const path = "/api/trans/vip/translate"

var errCodeMessage = map[string]string{
	"52001": "请求超时, 请稍后重试",
	"52002": "系统错误, 请稍后重试",
	"52003": "未授权用户, 检查您的 appid 是否正确，或者服务是否开通",
	"54000": "必填参数为空, 检查是否少传参数",
	"54001": "签名错误, 请检查您的签名生成方法",
	"54003": "访问频率受限, 请降低您的调用频率",
	"54004": "账户余额不足, 请前往管理控制台为账户充值",
	"54005": "长query请求频繁, 请降低长query的发送频率，3s后再试",
	"58000": "客户端IP非法",
}

// transRes is response body from remote api
type transRes struct {
	ErrorCode   string        `json:"error_code"`
	ErrorMsg    string        `json:"error_msg"`
	From        string        `json:"from"`
	To          string        `json:"to"`
	TransResult []TransResult `json:"trans_result"`
}

// TransResult is TransRes TransResult field
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

func wordsContainChinese(input string) bool {
	return utf8.RuneCountInString(input) != len(input)
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
func parseResponse(rc io.ReadCloser) (transRes, error) {
	var raw transRes

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
