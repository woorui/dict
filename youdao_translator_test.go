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
	}
	for _, table := range tables {
		res := genInput(table.p)
		if res != table.result {
			t.Errorf("genInput(%s) was incorrect, got:%s, want:%s", table.p, res, table.result)
		}
	}
}
