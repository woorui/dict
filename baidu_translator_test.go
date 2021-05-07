package main

import (
	"net/http"
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

func Test_baiduTranslator_genRequestURL(t *testing.T) {
	type fields struct {
		name    string
		client  *http.Client
		baseurl string
		appID   string
		secret  string
	}
	type args struct {
		text string
		salt string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "baidu",
			fields: fields{
				name:    "baidu",
				baseurl: "ss",
				appID:   "dsa",
				secret:  "dasdasd",
			},
			args: args{
				text: "adas",
				salt: "asdasd",
			},
			want: "ss?appid=dsa&from=auto&q=adas&salt=asdasd&sign=e9c551473b204d235a7010c6e0566b8f&to=zh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := &baiduTranslator{
				name:    tt.fields.name,
				client:  tt.fields.client,
				baseurl: tt.fields.baseurl,
				appID:   tt.fields.appID,
				secret:  tt.fields.secret,
			}
			if got := translator.genRequestURL(tt.args.text, tt.args.salt); got != tt.want {
				t.Errorf("baiduTranslator.genRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
