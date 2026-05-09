package agreementPay

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 协议支付-订单查询接口
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa
func AgreementQueryOrder(config *AgreementPayConfig, req *AgreementQueryOrderRequest) (map[string]string, error) {
	// 交易类型
	txnType := TRANS_TYPE_QUERY_ORDER

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	reqMap := map[string]string{
		"send_time":         timestamp,
		"msg_id":            req.MsgId,
		"version":           "4.0.0.0",
		"terminal_id":       config.TerminalId,
		"txn_type":          txnType,
		"member_id":         config.MemberId,
		"orig_trans_id":     req.OrigTransId,
		"orig_trade_date":   req.OrigTradeDate,
		"req_reserved1":     req.ReqReserved1,
		"req_reserved2":     req.ReqReserved2,
		"additional_info1":  req.AdditionalInfo1,
		"additional_info2":  req.AdditionalInfo2,
		"agent_member_id":   req.AgentMemberId,
		"agent_terminal_id": req.AgentTerminalId,
	}

	reqParams, err := generateRequestParams(config, reqMap)
	if err != nil {
		return nil, err
	}

	// 创建响应变量
	var responseStr string

	responseStr, err = sendRequest(
		config.GetBaseURL(),
		reqParams,
	)
	if err != nil {
		return nil, err
	}

	resp := getParams(responseStr)

	// 校验响应数据
	if err = verifyResponseData(config, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
