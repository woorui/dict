package main

// --- baidu api json struct below ---

// BaiduTranslateResult is response body from remote api
type BaiduTranslateResult struct {
	ErrorCode   string `json:"error_code"`
	ErrorMsg    string `json:"error_msg"`
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
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
	Web []YoudaoTranslateResultWeb `json:"web"`
}

// YoudaoTranslateResultWeb is sub-struct of YoudaoTranslateResult
type YoudaoTranslateResultWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}
