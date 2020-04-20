package main

import (
	"net/http"

	"github.com/gosuri/uitable"
)

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

// TranslatorFactory product some translator
// If you want add a new one, change code comments.
func TranslatorFactory(configs []Config, client *http.Client) *[]Translator {
	var arr []Translator
	for _, config := range configs {
		if config.Key == "baidu" {
			arr = append(arr, newBaiduTranslator(client, baidubURL, config.Value))
		}
		if config.Key == "youdao" {
			arr = append(arr, newYoudaoTranslator(client, baidubURL, config.Value))
		}
		// Add new translator here.
	}
	return &arr
}

// Translate translate the text with mutiple engine
func Translate(translators []Translator, text string) []Row {
	table := uitable.New()
	table.AddRow(title...)
	// TODO dorequest
	var rows []Row
	return rows
}
