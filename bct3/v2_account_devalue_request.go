package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 账户减值（分账减值）接口
// https://docs.baofu.com/docs/bct3/bct3-5002-001-01
func AccountDevalueRequest(config *BCT3Config, req *AccDevalueReq) (*AccDevalueResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_DEVALUE

	// 固定配置
	req.Version = "1.0.0"

	// 固定使用当前时间
	req.ReqTime = tool.GetReqTime()

	// 自动计算减值总金额
	req.CalcDealAmount()

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

	resp := &AccDevalueResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
