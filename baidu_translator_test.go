package main

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

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
			t.Errorf("generateHashSign(%s,%s,%s,%s) was incorrect, got:%s, want:%s", table.appid, table.q, table.salt, table.secret, res, table.sign)
		}
	}
}

func Test_genRequestURL(t *testing.T) {
	client := &http.Client{}
	tables := []struct {
		baseurl string
		appID   string
		secret  string
		text    string
		salt    string
		returns string
		err     error
	}{
		{"http://test.com/test/path", "mockAppid", "mockSign", "name", "123", "http://test.com/test/path?appid=mockAppid&from=auto&q=name&salt=123&sign=ee6dde35f3c80f27d892ab534ac0866e&to=zh", nil},
		{"://test.com/test/path", "mockAppid", "mockSign", "name", "321", "", errors.New("parse ://test.com/test/path: missing protocol scheme")},
	}
	for _, table := range tables {
		translator := newBaiduTranslator(client, table.baseurl, strings.Join([]string{table.appID, table.secret}, "-"))
		returns, err := translator.genRequestURL(table.text, table.salt)
		if err != nil {
			if err.Error() != table.err.Error() {
				t.Errorf(
					"translator(client, %s, %s, %s).genRequestURL(%s, %s) result err was incorrect, got:%s, want:%s",
					table.baseurl, table.appID, table.secret, table.text, table.salt, err.Error(), table.err.Error())
			}
		}
		if returns != table.returns {
			t.Errorf(
				"translator(client, %s, %s, %s).genRequestURL(%s, %s) result err was correct, got:%s, want:%s",
				table.baseurl, table.appID, table.secret, table.text, table.salt, returns, table.returns)
		}
	}
}
