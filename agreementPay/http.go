package agreementPay

import (
	"fmt"
	"log"
	"strings"

	"github.com/shushu2013/baofuSdk/tool"

	"github.com/pkg/errors"
)

func sendRequest(url string, params map[string]interface{}) (string, error) {

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	paramsStr, _ := tool.StringifyJSON(params)

	headersStr, _ := tool.StringifyJSON(headers)
	log.Printf("baofu_request_res url: %s, paramsStr: %s, headersStr: %s", url, paramsStr, headersStr)

	response, err := tool.SendPostHttpRequest(
		url,
		params,
		headers,
	)

	var respStr string

	defer func() {
		if err != nil {
			log.Printf("baofu_request_err url: %s, err: %s", url, err.Error())
		} else {
			log.Printf("baofu_request_res url: %s, body: %s", url, respStr)
		}
	}()

	if err != nil {
		tool.SendRobotWarning(
			fmt.Sprintf("宝付 API 调用报错, url: %s, params: %s", url, paramsStr),
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
			fmt.Sprintf("宝付 API 返回空串, url: %s, params: %s", url, paramsStr),
			err,
		)
	}

	return "", err
}

func getParams(response string) map[string]string {
	// 将 a=b&c=d 转换为 map[string]string
	respMap := make(map[string]string)
	pairs := strings.Split(response, "&")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			respMap[kv[0]] = kv[1]
		}
	}

	return respMap
}

// 从 KeyStr 中解析出 AES 密钥
func getAesKey(keyStr string) (string, error) {
	listKeyObj := strings.Split(keyStr, "|")
	if len(listKeyObj) == 2 {
		trimmed := strings.TrimSpace(listKeyObj[1])
		if trimmed != "" {
			return trimmed, nil
		} else {
			return "", errors.New("Key is Null!")
		}
	} else {
		return "", errors.New("Data format is incorrect!")
	}
}

// 生成接口请求参数
func generateRequestParams(config *AgreementPayConfig, reqMap map[string]string) (map[string]interface{}, error) {

	signStr := tool.CoverMap2String(reqMap)
	signature, err := tool.Sign(tool.Sha1X16(signStr), config.PrivateKey, tool.SIGNATURE_SHA1_WITH_RSA_ALGORITHM)
	if err != nil {
		return nil, err
	}

	reqMap["signature"] = signature

	params := make(map[string]interface{})
	for k, v := range reqMap {
		params[k] = v
	}

	return params, nil
}

// 校验接口返回值
func verifyResponseData(config *AgreementPayConfig, respMap map[string]string) error {

	if len(respMap) == 0 {
		err := errors.New("报文异常")
		return err
	}

	// 校验系统返回码
	if respMap["biz_resp_code"] != BIZ_RESP_CODE_SUCCESS {
		return errors.Errorf("%s:%s", respMap["biz_resp_code"], respMap["biz_resp_msg"])
	}

	// 校验商户接口应答码
	if respMap["resp_code"] == RESP_CODE_FAIL {
		return errors.Errorf("%s:%s", respMap["resp_code"], "应答码状态为失败")
	}

	// 校验签名
	signature := respMap["signature"]
	delete(respMap, "signature")

	if len(signature) == 0 {
		return errors.New("返回报文：缺少验签参数")
	}

	signStr := tool.Sha1X16(tool.CoverMap2String(respMap))
	if ok, _ := tool.Verify(signStr, config.PublicKey, signature, tool.SIGNATURE_SHA1_WITH_RSA_ALGORITHM); !ok {
		return errors.New("返回报文：验签失败")
	}

	return nil
}
