package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"

	"github.com/pkg/errors"
)

// 用户开户接口
// https://docs.baofu.com/docs/bct3/bct3-1g42bqv7t813j
func AccountOpenRequest(config *BCT3Config, req *AccOpenReq) (*AccOpenResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_OPEN
	// 创建AES密钥
	aesKey := tool.CreateAeskey(16)
	// 加密数字信封
	dgtlEnvlp, err := tool.EncryptByPublicKey(aesKey, config.PublicKey)
	if err != nil {
		return nil, err
	}

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 账户类型
	accType := req.AccType

	// 固定配置
	req.Version = "1.0.0"
	req.BusinessType = "BCT3"

	// 加密敏感数据
	req.AccInfo.CertificateNo, _ = tool.AesEncrypt(req.AccInfo.CertificateNo, aesKey)
	req.AccInfo.CardNo, _ = tool.AesEncrypt(req.AccInfo.CardNo, aesKey)
	if accType == ACCOUNT_TYPE_PERSONAL {
		req.AccInfo.MobileNo, _ = tool.AesEncrypt(req.AccInfo.MobileNo, aesKey)
	} else {
		req.AccInfo.CorporateCertId, _ = tool.AesEncrypt(req.AccInfo.CorporateCertId, aesKey)
		if req.AccInfo.ContactMobile != "" {
			req.AccInfo.ContactMobile, _ = tool.AesEncrypt(req.AccInfo.ContactMobile, aesKey)
		}
		if req.AccInfo.CorporateMobile != "" {
			req.AccInfo.CorporateMobile, _ = tool.AesEncrypt(req.AccInfo.CorporateMobile, aesKey)
		}
	}

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

	resp := &AccOpenResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	// 判断开户是否失败
	result := resp.Result[0]
	if result.State == STATE_FAILURE ||
		result.State == STATE_EXCEPTION {
		return nil, errors.Errorf("开户失败:%s:%s", result.ErrorCode, result.ErrorMsg)
	}

	return resp, nil
}
