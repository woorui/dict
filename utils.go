package main

import (
	"io/ioutil"
	"net/http"
	"unicode/utf8"
)

// HTTPGetRequest do a get request and return response body
func HTTPGetRequest(client *http.Client, url string) ([]byte, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func textContainChinese(text string) bool {
	return utf8.RuneCountInString(text) != len(text)
}
