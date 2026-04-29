package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"

	"github.com/pkg/errors"
)

// 账户余额查询接口
// https://docs.baofu.com/docs/bct3/bct3-2001-001-01
func AccountBalanceQueryRequest(config *BCT3Config, req *AccBalanceQueryReq) (*AccBalanceQueryResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_BALANCE_QUERY

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

	resp := &AccBalanceQueryResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	// 判断查询状态
	if resp.RetCode == RET_CODE_FAILURE {
		return nil, errors.Errorf("账户余额查询失败:%s", resp.ErrorMsg)
	}

	return resp, nil
}
