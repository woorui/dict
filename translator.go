package main

import (
	"net/http"
	"strings"

	"github.com/gosuri/uitable"
)

// Engine has multiple Translator
type Engine struct {
	Translators []Translator
}

// Translator has Translate method
type Translator interface {
	GetName() string
	Translate(text string) ([]Translation, error)
}

// Translation is the returning of Translator.Translate
type Translation struct {
	DataSource string
	Src        string
	Dst        string
	Phonetic   string
	Explain    string
}

// NewEngine construct some translator
// If you want add a new one, change code comments below.
func NewEngine(config []Config) *Engine {
	client := &http.Client{}
	var arr []Translator
	for _, c := range config {
		if c.Key == "youdao" {
			arr = append(arr, newYoudaoTranslator(client, baidubURL, c.Value))
		}
		if c.Key == "baidu" {
			arr = append(arr, newBaiduTranslator(client, baidubURL, c.Value))
		}
		// ADD HERE.
	}
	return &Engine{Translators: arr}
}

// Translate translate the str with mutiple engine,
// str will split by space in some sub string.
func (engine *Engine) Translate(str string) (*uitable.Table, error) {
	table := initTable()
	str = strings.Trim(str, " ")
	for _, translator := range engine.Translators {
		translations, err := translator.Translate(str)
		if err != nil {
			return table, err
		}
		for _, t := range translations {
			table.AddRow(t.DataSource, t.Src, t.Dst, t.Phonetic, t.Explain)
		}
	}
	return table, nil
}

func initTable() *uitable.Table {
	table := uitable.New()
	table.AddRow(tableTitle...)
	table.MaxColWidth = 60
	table.Wrap = true
	return table
}
