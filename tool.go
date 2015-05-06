package goflow

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/lunny/log"
	"github.com/satori/go.uuid"
)

//生成UUID
func NewUUID() string {
	return uuid.NewV4().String()
}

//字符串转整型
func StrToInt(value string) int {
	if value == "" {
		return 0
	}
	val, _ := strconv.Atoi(value)
	return val
}

//整型转字符串
func IntToStr(value int) string {
	return strconv.Itoa(value)
}

//装载XML文件
func LoadXML(xmlFile string) []byte {
	content, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		log.Errorf("error to read xml file %v", err)
		panic(fmt.Errorf("error to read xml file!"))
	}
	return content
}

//map转json
func MapToJson(v map[string]interface{}) string {
	if v == nil {
		return ""
	}
	js := simplejson.New()
	js.Set("map", v)
	ret, _ := js.Get("map").Encode()
	return string(ret)
}

//json转map
func JsonToMap(v string) map[string]interface{} {
	js, _ := simplejson.NewJson([]byte(v))
	return js.MustMap()
}

//删除Slice中的元素
func StringsRemove(strings []string, start, end int) []string {
	return append(strings[:start], strings[end:]...)
}

//删除Slice中的元素
func StringsRemoveAtIndex(strings []string, index int) []string {
	return StringsRemove(strings, index, index+1)
}

//格式化时间字符串
func FormatTime(t time.Time, layout string) string {
	if t.IsZero() {
		return ""
	} else {
		return t.Format(layout)
	}
}

func ProcessTime(args map[string]interface{}, timeParam string) time.Time {
	if timeParam != "" {
		var timeStr string
		if timeInf, ok := args[timeParam]; ok {
			timeStr = timeInf.(string)
		} else {
			timeStr = timeParam
		}
		the_time, err := time.Parse(STD_TIME_LAYOUT, timeStr)
		if err == nil {
			return the_time
		}
	}
	return time.Time{}
}
