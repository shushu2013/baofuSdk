package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func SendPostHttpRequest(urlStr string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	return sendPostHttpRequest(urlStr, params, headers, 20)
}

func sendPostHttpRequest(urlStr string, params map[string]interface{}, headers map[string]string, timeout time.Duration) (*http.Response, error) {

	startTime := time.Now()

	var reader io.Reader

	contentType := strings.ToLower(headers["Content-Type"])

	if params != nil {
		if contentType == "application/x-www-form-urlencoded" {
			values := url.Values{}

			for key := range params {
				if valueStr, ok := params[key].(string); ok {
					values.Set(key, valueStr)
				}
			}
			reader = bytes.NewBufferString(values.Encode())
		} else {
			value, err := json.Marshal(params)
			if err != nil {
				return nil, err
			}
			//fmt.Println("debug params", string(value))
			reader = bytes.NewBuffer(value)
		}
	}

	req, err := http.NewRequest("POST", urlStr, reader)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	client := http.Client{Timeout: timeout * time.Second}
	resp, err := client.Do(req)

	duration := time.Now().Sub(startTime).Milliseconds()
	//fmt.Printf("debug request url: %s, duration: %dms\n", urlStr, duration)
	if duration > 3000 {
		SendRobotWarning(
			fmt.Sprintf("post url: %s, duration: %dms", urlStr, duration),
			errors.New("请求外部接口耗时过长"),
		)
	}

	return resp, err

}

func StringifyHttpResponse(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ToJsonResponse(response string, v interface{}) error {
	return ParseJSON(response, v)
}
