package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/pkg/errors"
)

func interfaceIsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

// https://cloud.tencent.com/developer/article/1661736
func ParseJSON(value string, v interface{}) error {
	if len(value) == 0 {
		return errors.New("json is empty")
	}
	if err := json.Unmarshal([]byte(value), v); err != nil {
		log.Println(value, DumpStacks())
		return errors.Errorf("JSON 解析错误: 【%s】, err: %s", value, err)
	}
	return nil
}

func StringifyJSON(v interface{}) (string, error) {
	value, err := json.Marshal(v)
	if err != nil {
		fmt.Println("debug StringifyJSON", err)
		return "", err
	}
	// 如果 v 是一个空切片，则 value 序列化出来是 "null"
	result := string(value)
	if result == "null" {
		if interfaceIsSlice(v) {
			return "[]", nil
		}
		return "", nil
	}
	return result, nil
}

func PrintJSON(title string, v interface{}) {
	str, _ := StringifyJSON(v)
	log.Println(title, str)
}

func StringifyJSONWithoutEscape(v interface{}) (string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	if err != nil {
		fmt.Println("debug StringifyJSON", err)
		return "", err
	}
	result := string(buffer.Bytes())
	return result, nil
}
