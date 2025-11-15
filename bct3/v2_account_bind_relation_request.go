package bct3

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 账户绑定关系操作接口
// https://docs.baofu.com/docs/bct3/bct3-1gahl9mq80c50
func AccountBindRelationRequest(config *BCT3Config, req *AccBindRelationReq) (*AccBindRelationResp, error) {
	// 服务编号
	serviceTp := SERVICE_ACCOUNT_BIND_RELATION_OPERATION

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 固定配置
	req.Version = "4.1.0"

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

	resp := &AccBindRelationResp{}
	if err = tool.ParseJSON(response.Body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
