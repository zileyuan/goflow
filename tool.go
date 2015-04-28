package goflow

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/lunny/log"
)

func StrToInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func LoadXML(xmlFile string) []byte {
	content, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		log.Errorf("error to read xml file %v", err)
		panic(fmt.Errorf("error to read xml file!"))
	}
	return content
}
