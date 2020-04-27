package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"unicode/utf8"

	"gopkg.in/yaml.v2"
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

// getConfig get config needs by engine from path p
func getConfig(p string) ([]Config, error) {
	var err error
	var b []byte
	var config []Config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}
	homep := filepath.Join(homeDir, configfile)
	if p == "" {
		if _, err = os.Stat(homep); err != nil {
			if os.IsNotExist(err) {
				return config, errFileIsRequired
			}
			return config, err
		}
		b, err = ioutil.ReadFile(homep)
		if err != nil {
			return config, err
		}
	} else {
		b, err = ioutil.ReadFile(p)
		if err != nil {
			return config, err
		}
		f, err := os.OpenFile(homep, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return config, err
		}
		defer f.Close()
		if _, err := f.Write(b); err != nil {
			return config, err
		}
		if err := f.Sync(); err != nil {
			return config, err
		}
	}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func readConfigFromFile(p string) ([]Config, error) {
	var config []Config
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func writeToFile(p string, b []byte) error {
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(b); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
