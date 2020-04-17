package main

// Translator has Translate method
type Translator interface {
	Name() string
	Translate(text string) Translation
}

// Translation is the returning of Translator.Translate
type Translation struct {
	Src      string
	Dst      string
	Phonetic string
	Explain  string
}
