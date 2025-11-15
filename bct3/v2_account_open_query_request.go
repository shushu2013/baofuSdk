package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 开户结果查询接口
// https://docs.baofu.com/docs/bct3/bct3-1g42c3304das0
func AccountOpenQueryRequest(config *BCT3Config, req *AccOpenQueryReq) (*AccOpenQueryResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_OPEN_QUERY

	// 创建AES密钥
	aesKey := tool.CreateAeskey(16)

	// 加密数字信封
	dgtlEnvlp, err := tool.EncryptByPublicKey(aesKey, config.PublicKey)
	if err != nil {
		return nil, err
	}

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 固定配置
	req.Version = "1.0.0"

	reqHeader := &RequestHeader{
		MemberID:   config.MemberId,
		TerminalID: config.TerminalId,
		Timestamp:  timestamp,
		VerifyType: config.VerifyType,
		Charset:    config.Charset,
		Version:    config.Version,
		SignSN:     config.SignSN,
		NcrptnSN:   config.NcrptnSN,
		DgtlEnvlp:  dgtlEnvlp,
	}

	reqParams := generateBCT3RequestParams(config, reqHeader, req)
	response := &ResponseData{}

	// 发送请求
	err = sendRequest(
		config.GetBaseURL(serviceTp),
		reqParams,
		response,
	)
	if err != nil {
		return nil, err
	}

	if err = verifyBCT3ResponseData(config, response); err != nil {
		return nil, err
	}

	responseHeader := &ResponseHeader{}
	if err = tool.ParseJSON(response.Header, responseHeader); err != nil {
		return nil, err
	}

	resp := &AccOpenQueryResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	// 解密敏感数据
	if resp.RetCode == RET_CODE_SUCCESS {
		resultAesk, err := tool.DecryptByPrivateKey(responseHeader.DgtlEnvlp, config.PrivateKey)
		if err != nil {
			return nil, err
		}

		// 绑定手机号
		if resp.Result.BindMobile != "" {
			resp.Result.BindMobile, err = tool.AesDecrypt(resp.Result.BindMobile, resultAesk)
			if err != nil {
				return nil, err
			}
		}

		// 证件号码（社会信用代码）
		if resp.Result.CertificateNo != "" {
			resp.Result.CertificateNo, err = tool.AesDecrypt(resp.Result.CertificateNo, resultAesk)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, nil
}
