package main

import (
	"io/ioutil"
	"strings"
)

func sed(old string, new string, file string) {
	filePath := file
	fileData, _ := ioutil.ReadFile(filePath)
	fileString := string(fileData)
	fileString = strings.ReplaceAll(fileString, old, new)
	fileData = []byte(fileString)
	_ = ioutil.WriteFile(filePath, fileData, 0o600)
}
