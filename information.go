package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const saveInformationFileName = ".dict"

// homedir is user's home directory, For example It is /home/$USER in linux environment
var homeDir string

func init() {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(errors.New("Can not get homeDir" + err.Error()))
	}
	homeDir = usr.HomeDir
}

// Information was applied from baidu open api
// apply link: http://api.fanyi.baidu.com/api/trans/product/index
type Information struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

// readFromFile read SomeThing from SomeFile
func readFromFile(filePath string) ([]byte, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, errors.New("filepath " + filePath + " does not exist")
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// getInformation get information that cantain appid and secret from ~/.dict
func getInformation() (*Information, error) {
	var info Information
	b, err := readFromFile(filepath.Join(homeDir, saveInformationFileName))
	if err != nil {
		return &info, err
	}
	if err := json.Unmarshal(b, &info); err != nil {
		return &info, err
	}
	return &info, err
}

// removeFile delete specify file
func removeFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func removeInformation() {
	fpd := filepath.Join(homeDir, saveInformationFileName)
	if err := removeFile(fpd); err != nil {
		log.Fatal("delete file " + fpd + " failed")
	}
}

// writeToFile save SomeThing to SomeFile
func writeToFile(filePath string, content []byte) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := f.Write(content); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}

// saveInformation save information that cantain appid and secret to ~/.dict
func saveInformation(info *Information) error {
	fp := filepath.Join(homeDir, saveInformationFileName)

	j, err := json.Marshal(&info)
	if err != nil {
		return err
	}
	if err := writeToFile(fp, j); err != nil {
		return err
	}
	return nil
}
