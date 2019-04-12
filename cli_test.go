package main

import (
	"bufio"
	"strings"
	"testing"
)

func Test_getTrimInput(t *testing.T) {
	tables := []struct {
		reader     *bufio.Reader
		allowEmpty bool
		result     string
		err        error
	}{
		{bufio.NewReader(strings.NewReader("abc\n")), true, "abc", nil},
		{bufio.NewReader(strings.NewReader("abc")), true, "", nil},
		{bufio.NewReader(strings.NewReader("\n")), true, "", nil},
		// {bufio.NewReader(strings.NewReader("123\n")), false, "", nil},
	}
	for _, table := range tables {
		res, err := getTrimInput(table.reader, table.allowEmpty)
		if err != nil && table.err != nil {
			if err.Error() != table.err.Error() {
				t.Errorf("getTrimInput err value incorrent")
			}
		}
		if res != table.result {
			t.Errorf("getTrimInput(reader,%t) was incorrect, got:%s, want:%s", table.allowEmpty, res, table.result)
		}
	}
}

func Test_formatPes(t *testing.T) {
	strings.NewReader("abc")
	tables := []struct {
		tr  []TransResult
		res string
	}{
		{[]TransResult{{Src: "word", Dst: "单词"}}, "|> word: 单词"},
		{[]TransResult{{Src: "单词", Dst: "word"}}, "|> 单词: word"},
		{[]TransResult{}, ""},
	}
	for _, table := range tables {
		res := formatPes(table.tr)
		if res != table.res {
			t.Errorf("formatPes(%v) was incorrect, got:%s, want:%s", table.tr, res, table.res)
		}
	}
}
