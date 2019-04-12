package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"reflect"
	"testing"
)

func genMockResopnse(res transRes) io.ReadCloser {
	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln("Mock response failed, please check args later")
	}
	return ioutil.NopCloser(bytes.NewReader(b))
}

// compareTransRes return two transRes is equal or not
func compareTransRes(left transRes, right transRes) (isEq bool) {
	if left.ErrorCode != right.ErrorCode || left.ErrorMsg != right.ErrorMsg || left.From != right.From || left.To != right.To {
		return false
	}
	var lrs, rrs = left.TransResult, right.TransResult
	return reflect.DeepEqual(lrs, rrs)
}

func Test_generateHashSign(t *testing.T) {
	tables := []struct {
		appid  string
		q      string
		salt   string
		secret string
		sign   string
	}{
		{"aaaaaa", "bbbbb", "cccccc", "dddddd", "858dc648eda3116267cc410e34f521ff"},
		{"eeeeee", "fffff", "gggggg", "hhhhhh", "cf50e6b7140da6d1cf390df13e17baf1"},
		{"iiiiii", "jjjjjj", "kkkkkk", "dddddd", "c7ec9a743d9072db79b0624cc804c55b"},
	}
	for _, table := range tables {
		res := generateHashSign(table.appid, table.q, table.salt, table.secret)
		if res != table.sign {
			t.Errorf("wordsContainChinese(%s,%s,%s,%s) was incorrect, got:%s, want:%s", table.appid, table.q, table.salt, table.secret, res, table.sign)
		}
	}
}

func Test_wordsContainChinese(t *testing.T) {
	tables := []struct {
		word           string
		containChinese bool
	}{
		{"I love You", false},
		{"我 love You", true},
		{"我爱你", true},
	}
	for _, table := range tables {
		res := wordsContainChinese(table.word)
		if res != table.containChinese {
			t.Errorf("wordsContainChinese(%s) was incorrect, got:%v, want:%v", table.word, res, table.containChinese)
		}
	}
}

func Test_genRequestURL(t *testing.T) {
	tables := []struct {
		baseurl  string
		path     string
		word     string
		appid    string
		salt     string
		sign     string
		parseurl string
		err      error
	}{
		{"http://test.com", "/test/path", "name", "mockAppid", "112233", "mockSign",
			"http://test.com/test/path?appid=mockAppid&from=auto&q=name&salt=112233&sign=mockSign&to=zh", nil},
		{"://test.com", "/test/path", "name", "mockAppid", "112233", "mockSign",
			"", errors.New("parse ://test.com: missing protocol scheme")},
	}
	for _, table := range tables {
		url, err := genRequestURL(table.baseurl, table.path, table.word, table.appid, table.salt, table.sign)
		if err != nil {
			if err.Error() != table.err.Error() {
				t.Errorf(
					"genRequestURL(%s,%s,%s,%s,%s,%s) result err was incorrect, got:%s, want:%s",
					table.baseurl, table.path, table.word, table.appid, table.salt, table.sign, err.Error(), table.err.Error())
			}
		}
		if url != table.parseurl {
			t.Errorf(
				"genRequestURL(%s,%s,%s,%s,%s,%s) result url was incorrect, got:%s, want:%s",
				table.baseurl, table.path, table.word, table.appid, table.salt, table.sign, url, table.parseurl)
		}
	}
}

func Test_parseResponse(t *testing.T) {
	mockTransRes := transRes{
		ErrorCode: "52001",
		ErrorMsg:  "链接超时",
		From:      "en",
		To:        "zh",
		TransResult: []TransResult{
			{Src: "hello", Dst: "你好"},
		},
	}
	tables := []struct {
		body io.ReadCloser
		res  transRes
		err  error
	}{
		{genMockResopnse(mockTransRes), mockTransRes, nil},
	}
	for _, table := range tables {
		rs, err := parseResponse(table.body)
		if err != nil {
			if table.err != nil {
				if err.Error() != table.err.Error() {
					t.Error("parseResponse() return error value incorrent")
				}
				t.Error("parseResponse() return error nil, but except not nil")
			}
		}
		if !compareTransRes(rs, table.res) {
			t.Errorf("parseResponse() result was incorrect, got:%s, %s, want:%s, %s", rs, err.Error(), table.res, table.err.Error())
		}
	}
}
