package main

import "testing"

func Test_textContainChinese(t *testing.T) {
	tables := []struct {
		word           string
		containChinese bool
	}{
		{"I love You", false},
		{"我 love You", true},
		{"我爱你", true},
	}
	for _, table := range tables {
		res := textContainChinese(table.word)
		if res != table.containChinese {
			t.Errorf("textContainChinese(%s) was incorrect, got:%v, want:%v", table.word, res, table.containChinese)
		}
	}
}
