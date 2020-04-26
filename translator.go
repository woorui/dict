package main

import (
	"fmt"
	"net/http"
	"strings"
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

// Translate translate the str with mutiple engine,
// str will split by space in some sub string.
func (engine *Engine) Translate(str string) error {
	texts := strings.Split(str, " ")
	tch := make(chan Translation)
	errch := make(chan error, 1)
	for _, translator := range engine.Translators {
		for _, text := range texts {
			go func(text string) {
				translations, err := translator.Translate(text)
				if err != nil {
					errch <- err
				}
				for _, t := range translations {
					tch <- t
				}
			}(text)
		}
	}

	for {
		select {
		case t := <-tch:
			fmt.Println("--", t)
		case err := <-errch:
			return err
		}
	}
	return nil
}
