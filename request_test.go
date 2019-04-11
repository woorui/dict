package main

import "testing"

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
