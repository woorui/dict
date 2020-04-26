package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

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
		if c.Key == "baidu" {
			arr = append(arr, newBaiduTranslator(client, baidubURL, c.Value))
		}
		if c.Key == "youdao" {
			arr = append(arr, newYoudaoTranslator(client, baidubURL, c.Value))
		}
		// ADD HERE.
	}
	return &Engine{Translators: arr}
}

// Translate translate the text with mutiple engine
func (engine Engine) Translate(str string, done chan struct{}) (*uitable.Table, error) {
	var wg sync.WaitGroup
	texts := strings.Split(str, " ")
	errch := make(chan error, 1)
	q := make(chan Translation)
	for _, translator := range engine.Translators {
		for _, text := range texts {
			wg.Add(1)
			go func(translator Translator, text string) {
				translation, err := translator.Translate(text)
				if err != nil {
					errch <- err
				}
				for _, t := range translation {
					q <- t
				}
				wg.Done()
			}(translator, text)
		}
	}
	wg.Wait()
	done <- struct{}{}

	table := &uitable.Table{}
	select {
	case t := <-q:
		table.AddRow(t.DataSource, t.Src, t.Dst, t.Phonetic, t.Explain)
	case <-done:
		break
	case err := <-errch:
		return table, err
	}
	return table, nil
}

func initTable() *uitable.Table {
	table := uitable.New()
	table.AddRow(title...)
	return table
}

func subscriber(q chan Translation, done chan struct{}, errch chan error) {
	table := &uitable.Table{}
	select {
	case t := <-q:
		table.AddRow(t.DataSource, t.Src, t.Dst, t.Phonetic, t.Explain)
	case <-done:
		fmt.Println(table)
		table = initTable()
	case err := <-errch:
		fmt.Println(err.Error())
		break
	}
}

// Run run app
func Run() {
	configs := []Config{}
	engine := NewEngine(configs)
	done := make(chan struct{}, 1)
	engine.Translate("", done)
}
