package main

import (
	"net/http"
	"testing"
)

func Test_genInput(t *testing.T) {
	tables := []struct {
		p      string
		result string
	}{
		{"asdfghjklasdfghjklasdfghjkl", "asdfghjkla27lasdfghjkl"},
		{"asdfghjkl", "asdfghjkl"},
		{"aaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaa"},
		{"香蕉苹果梨橘子香蕉苹果梨橘子香蕉苹果梨橘子", "香蕉苹果梨橘子香蕉苹21梨橘子香蕉苹果梨橘子"},
		{"栗子", "栗子"},
		{"t push up middle key", "t push up middle key"},
	}
	for _, table := range tables {
		res := genInput(table.p)
		if res != table.result {
			t.Errorf("genInput(%s) was incorrect, got:%s, want:%s", table.p, res, table.result)
		}
	}
}

func Test_youdaoTranslator_genRequestURL(t *testing.T) {
	type fields struct {
		name      string
		client    *http.Client
		baseurl   string
		appKey    string
		appSecret string
	}
	type args struct {
		text      string
		timestamp string
		salt      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "youdao",
			fields: fields{
				appKey:    "111",
				appSecret: "222",
			},
			args: args{
				text:      "112",
				timestamp: "22",
				salt:      "wwwww",
			},
			want: "https://openapi.youdao.com/api?appKey=111&curtime=22&from=en&q=112&salt=wwwww&sign=c36e79f75c266117a51c78a389d17b4815b3db3284038fbcf5a4bb9c9e5242b5&signType=v3&to=zh-CHS",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			translator := &youdaoTranslator{
				name:      tt.fields.name,
				client:    tt.fields.client,
				baseurl:   tt.fields.baseurl,
				appKey:    tt.fields.appKey,
				appSecret: tt.fields.appSecret,
			}
			if got := translator.genRequestURL(tt.args.text, tt.args.timestamp, tt.args.salt); got != tt.want {
				t.Errorf("youdaoTranslator.genRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
