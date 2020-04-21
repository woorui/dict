package main

import (
	"context"
	"net/http"
	"sync"

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
func Translate(translators []Translator, text string) *uitable.Table {
	table := uitable.New()
	table.AddRow(title...)
	for _, translator := range translators {
		translation, err := translator.Translate(text)
		if err != nil {
			panic(err)
		}
		for _, t := range translation {
			table.AddRow(t.DataSource, t.Src, t.Dst, t.Phonetic, t.Explain)
		}
	}
	return table
}

func translate(translators []Translator, text string, q chan Translation) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, translator := range translators {
		go func(translator Translator) {
			translation, err := translator.Translate(text)
			if err != nil {
				err = ctx.Err()
			}
			for _, t := range translation {
				q <- t
			}
			cancel()
		}(translator)
	}
}

func translateWaitGroup(translators []Translator, text string) {
	var wg sync.WaitGroup
	q := make(chan Translation)
	for _, translator := range translators {
		wg.Add(len(translators))
		go func(translator Translator) {
			translation, err := translator.Translate(text)
			if err != nil {
				panic(err)
			}
			for _, t := range translation {
				q <- t
			}
			wg.Done()
		}(translator)
	}
	wg.Wait()
	close(q)

}
