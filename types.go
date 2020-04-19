package main

// Translator has Translate method
type Translator interface {
	GetName() string
	Translate(text string) Translation
}

// Translation is the returning of Translator.Translate
type Translation struct {
	Src      string
	Dst      string
	Phonetic string
	Explain  string
}

// --- baidu api json struct below ---

// BaiduTranslateResult is response body from remote api
type BaiduTranslateResult struct {
	ErrorCode   string             `json:"error_code"`
	ErrorMsg    string             `json:"error_msg"`
	From        string             `json:"from"`
	To          string             `json:"to"`
	TransResult []BaiduTransResult `json:"trans_result"`
}

// BaiduTransResult is type of BaiduTranslateResult.TransResult
type BaiduTransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

// --- youdao json struct below ---

// YoudaoTranslateResult is response of youdao-api
type YoudaoTranslateResult struct {
	ErrorCode   string   `json:"errorCode"`
	Translation []string `json:"translation"`
	Query       string   `json:"query"`
	Basic       struct {
		Phonetic string   `json:"phonetic"`
		Explains []string `json:"explains"`
	} `json:"basic"`
	Web     []YoudaoTranslateResultWeb `json:"web"`
	Webdict string                     `json:"webdict"`
}

// YoudaoTranslateResultWeb is sub-struct of YoudaoTranslateResult
type YoudaoTranslateResultWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}
