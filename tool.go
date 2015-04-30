package goflow

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/bitly/go-simplejson"
	"github.com/lunny/log"
	"github.com/satori/go.uuid"
)

func NewUUID() string {
	return uuid.NewV4().String()
}

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

func MapToJson(v map[string]interface{}) string {
	js := simplejson.New()
	js.Set("map", v)
	ret, _ := js.Get("map").String()
	return ret
}

func JsonToMap(v string) map[string]interface{} {
	js, _ := simplejson.NewJson([]byte(v))
	return js.MustMap()
}

func StringsRemove(strings []string, start, end int) []string {
	return append(strings[:start], strings[end:]...)
}

func StringsRemoveAtIndex(strings []string, index int) []string {
	return StringsRemove(strings, index, index+1)
}
