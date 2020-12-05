package utils

import (
	"io/ioutil"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInput(path string, seperator string) []string {
	rawInput, err := ioutil.ReadFile(path)
	CheckError(err)
	return strings.Split(string(rawInput), seperator)
}
