package refund

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/shushu2013/baofuSdk/tool"

	"github.com/pkg/errors"
)

// sendRequest 发送HTTP请求
func sendRequest(url string, params map[string]interface{}) (string, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	paramsStr, _ := tool.StringifyJSON(params)
	headersStr, _ := tool.StringifyJSON(headers)
	log.Printf("baofu_refund_request url: %s, paramsStr: %s, headersStr: %s", url, paramsStr, headersStr)

	response, err := tool.SendPostHttpRequest(
		url,
		params,
		headers,
	)

	var respStr string

	defer func() {
		if err != nil {
			log.Printf("baofu_refund_request_err url: %s, err: %s", url, err.Error())
		} else {
			log.Printf("baofu_refund_request_res url: %s, body: %s", url, respStr)
		}
	}()

	if err != nil {
		tool.SendRobotWarning(
			fmt.Sprintf("宝付退款相关API调用报错, url: %s, params: %s", url, paramsStr),
			err,
		)
		err = errors.New("系统开小差了，请重试哦")
		return "", err
	}

	if respStr, err = tool.StringifyHttpResponse(response); err != nil {
		return "", err
	}
	if len(respStr) > 0 {
		return respStr, nil
	} else {
		err = errors.New("系统异常")
		tool.SendRobotWarning(
			fmt.Sprintf("宝付退款API返回空串, url: %s, params: %s", url, paramsStr),
			err,
		)
	}

	return "", err
}

// parseAndDecryptResponse 解析并解密响应数据
// 1. 使用宝付公钥解密响应字符串
// 2. Base64解码
// 3. JSON解析为结构体
func parseAndDecryptResponse(config *RefundConfig, responseStr string, result interface{}) error {
	// 1. RSA解密（使用宝付公钥）
	decryptedStr, err := tool.DecryptByPublicKey(responseStr, config.PublicKey)
	if err != nil {
		return errors.Wrap(err, "响应数据解密失败")
	}

	// 2. Base64解码
	decryptedBytes, err := base64.StdEncoding.DecodeString(decryptedStr)
	if err != nil {
		return errors.Wrap(err, "Base64解码失败")
	}

	// 3. JSON解析为结构体
	if err := tool.ParseJSON(string(decryptedBytes), result); err != nil {
		return errors.Wrap(err, "JSON解析失败")
	}

	return nil
}
