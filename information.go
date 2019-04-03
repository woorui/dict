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

const saveInformationFileName = ".simple_dict"

// Information is applied from baidu open api
// apply link: http://api.fanyi.baidu.com/api/trans/product/index
type Information struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

// homedir is user's home directory, For example It is /home/$USER in linux environment
var homeDir = (func() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalln(errors.New("Can not get homeDir" + err.Error()))
	}
	return usr.HomeDir
})()

func getInformation() (*Information, error) {
	var info Information
	fpd := filepath.Join(homeDir, saveInformationFileName)
	if _, err := os.Stat(fpd); os.IsNotExist(err) {
		return &info, errors.New("file named " + saveInformationFileName + " does not exist")
	}
	b, err := ioutil.ReadFile(fpd)
	if err := json.Unmarshal(b, &info); err != nil {
		return &info, err
	}
	return &info, err
}

func removeInformation() {
	fpd := filepath.Join(homeDir, saveInformationFileName)
	if err := os.Remove(fpd); err != nil {
		log.Fatalln("delete "+fpd, " failed")
	}
}

func saveInformation(info *Information) error {
	fp := filepath.Join(homeDir, saveInformationFileName)

	j, err := json.Marshal(&info)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(j); err != nil {
		return err
	}
	if err := f.Sync(); err != nil {
		return err
	}
	return nil
}
