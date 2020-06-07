package main

import (
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
