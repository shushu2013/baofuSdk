package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 账户信息查询接口
// https://docs.baofu.com/docs/bct3/bct3-1g42ctm4vntec
func AccountInfoQueryRequest(config *BCT3Config, req *AccInfoQueryReq) (*AccInfoQueryResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_QUERY

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

	resp := &AccInfoQueryResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	// 解密敏感数据
	if resp.RetCode == RET_CODE_SUCCESS {
		// 解密响应头中的数字信封
		envelopeKey, err := tool.DecryptByPrivateKey(responseHeader.DgtlEnvlp, config.PrivateKey)
		if err != nil {
			return nil, err
		}

		// 解密账户信息中的敏感数据
		if resp.AccInfo.CertificateNo != "" {
			resp.AccInfo.CertificateNo, _ = tool.AesDecrypt(resp.AccInfo.CertificateNo, envelopeKey)
		}

		// 解密绑卡信息中的敏感数据
		for _, bindCardInfo := range resp.BindCardInfoList {
			if bindCardInfo.CardNo != "" {
				bindCardInfo.CardNo, _ = tool.AesDecrypt(bindCardInfo.CardNo, envelopeKey)
			}
			if bindCardInfo.MobileNo != "" {
				bindCardInfo.MobileNo, _ = tool.AesDecrypt(bindCardInfo.MobileNo, envelopeKey)
			}
		}
	}

	return resp, nil
}
