package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 账户升级接口
// https://docs.baofu.com/docs/bct3/bct3-1gahkbffhgche
func AccountUpgradeRequest(config *BCT3Config, req *AccUpgradeReq) (*AccUpgradeResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_UPGRADE
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
	req.Version = "4.1.0"

	// 加密敏感数据
	if req.ContactMobile != "" {
		req.ContactMobile, _ = tool.AesEncrypt(req.ContactMobile, aesKey)
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

	resp := &AccUpgradeResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
