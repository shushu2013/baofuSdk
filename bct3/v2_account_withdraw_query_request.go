package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 转账结果查询（取现）接口
// https://docs.baofu.com/docs/bct3/bct3-3005-001-01
func AccountWithdrawQueryRequest(config *BCT3Config, req *AccWithdrawQueryReq) (*AccWithdrawQueryResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_WITHDRAW_QUERY

	// 固定配置
	req.Version = "1.0.0"

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	reqHeader := &RequestHeader{
		MemberID:   config.MemberId,
		TerminalID: config.TerminalId,
		Timestamp:  timestamp,
		VerifyType: config.VerifyType,
		Charset:    config.Charset,
		Version:    config.Version,
		SignSN:     config.SignSN,
		NcrptnSN:   config.NcrptnSN,
	}

	reqParams := generateBCT3RequestParams(config, reqHeader, req)
	response := &ResponseData{}

	// 发送请求
	err := sendRequest(
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

	resp := &AccWithdrawQueryResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
