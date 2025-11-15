package bct3

import (
	"fmt"
	"log"

	"github.com/shushu2013/baofuSdk/tool"

	"github.com/pkg/errors"
)

func sendRequest(url string, params map[string]interface{}, result interface{}) error {

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
		return err
	}

	if respStr, err = tool.StringifyHttpResponse(response); err != nil {
		return err
	}
	if len(respStr) > 0 {
		if err = tool.ToJsonResponse(respStr, result); err != nil {
			return err
		}
	} else {
		err = errors.New("系统异常")
		tool.SendRobotWarning(
			fmt.Sprintf("宝付 API 返回空串, url: %s, params: %s", url, paramsStr),
			err,
		)
	}

	return err
}

// 生成宝财通3 接口请求参数
func generateBCT3RequestParams(config *BCT3Config, reqHeader *RequestHeader, req interface{}) map[string]interface{} {

	headerStr, _ := tool.StringifyJSON(reqHeader)
	bodyStr, _ := tool.StringifyJSON(req)

	signStr, _ := tool.Sign(headerStr+bodyStr, config.PrivateKey, tool.SIGNATURE_SHA256_WITH_RSA_ALGORITHM)

	reqData := map[string]interface{}{
		"header": headerStr,
		"body":   bodyStr,
		"sign":   signStr,
	}

	return reqData
}

// 校验宝财通3 接口返回值
func verifyBCT3ResponseData(config *BCT3Config, response *ResponseData) error {

	if len(response.Header) == 0 {
		err := errors.New("报文异常")
		return err
	}

	responseHeader := &ResponseHeader{}
	if err := tool.ParseJSON(response.Header, responseHeader); err != nil {
		return err
	}

	// 校验系统返回码
	if responseHeader.SysRespCode != SYS_RESP_CODE_SUCCESS {
		return errors.Errorf("%s:%s", responseHeader.SysRespCode, responseHeader.SysRespDesc)
	}

	// 校验签名
	if ok, _ := tool.Verify(response.Header+response.Body, config.PublicKey, response.Sign, tool.SIGNATURE_SHA256_WITH_RSA_ALGORITHM); !ok {
		return errors.New("返回报文：验签失败")
	}

	// 校验返回值
	baseResp := &AccBaseResp{}
	if err := tool.ParseJSON(response.Body, baseResp); err != nil {
		return err
	}
	// 校验业务返回码
	if baseResp.RetCode == RET_CODE_FAILURE {
		return errors.Errorf("%s:%s", baseResp.ErrorCode, baseResp.ErrorMsg)
	}

	return nil
}
