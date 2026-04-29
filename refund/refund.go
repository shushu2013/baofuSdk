package refund

import (
	"encoding/base64"

	"github.com/shushu2013/baofuSdk/tool"
)

// Refund 退款接口
// https://docs.baofu.com/docs/interface_document/RefundAPI
func (c *RefundClient) Refund(req *RefundRequest) (*RefundResponse, error) {
	// 构建加密数据
	encryptedData := &RefundEncryptedData{
		TerminalId:             c.config.TerminalId,
		MemberId:               c.config.MemberId,
		TxnSubType:             TXN_SUB_TYPE_REFUND,
		RefundType:             req.RefundType,
		TransId:                req.TransId,
		RefundOrderNo:          req.RefundOrderNo,
		TransSerialNo:          req.TransSerialNo,
		RefundReason:           req.RefundReason,
		RefundAmt:              req.RefundAmt,
		RefundTime:             req.RefundTime,
		AdditionalInfo:         req.AdditionalInfo,
		ReqReserved:            req.ReqReserved,
		NoticeUrl:              req.NoticeUrl,
		UnionRefundInfo:        req.UnionRefundInfo,
		PartShareRefundInfo:    req.PartShareRefundInfo,
		ShareRefundAdvanceFlag: req.ShareRefundAdvanceFlag,
		ShareRefundInfo:        req.ShareRefundInfo,
	}

	// 将加密数据转为JSON
	dataContentJSON, err := tool.StringifyJSON(encryptedData)
	if err != nil {
		return nil, err
	}

	// 先Base64编码，再使用商户私钥证书加密
	dataContentBase64 := base64.StdEncoding.EncodeToString([]byte(dataContentJSON))
	dataContentEncrypted, err := tool.RsaEncryptByPrivateKey(dataContentBase64, c.config.PrivateKey)
	if err != nil {
		return nil, err
	}

	// 构建请求参数
	reqParams := map[string]interface{}{
		"version":      "4.0.0.0",
		"member_id":    c.config.MemberId,
		"terminal_id":  c.config.TerminalId,
		"txn_type":     TXN_TYPE_REFUND,
		"txn_sub_type": TXN_SUB_TYPE_REFUND,
		"data_type":    "json",
		"data_content": dataContentEncrypted,
	}

	// 发送请求
	responseStr, err := sendRequest(
		c.config.GetBaseURL(),
		reqParams,
	)
	if err != nil {
		return nil, err
	}

	// 解析并解密响应数据
	resp := &RefundResponse{}
	if err := parseAndDecryptResponse(c.config, responseStr, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// RefundQuery 退款查询接口
// https://docs.baofu.com/docs/interface_document/RefundAPI
func (c *RefundClient) RefundQuery(req *RefundQueryRequest) (*RefundQueryResponse, error) {
	// 构建加密数据
	encryptedData := &RefundQueryEncryptedData{
		TxnSubType:     TXN_SUB_TYPE_REFUND_QUERY,
		MemberId:       c.config.MemberId,
		TerminalId:     c.config.TerminalId,
		RefundOrderNo:  req.RefundOrderNo,
		TransSerialNo:  req.TransSerialNo,
		AdditionalInfo: req.AdditionalInfo,
		ReqReserved:    req.ReqReserved,
	}

	// 将加密数据转为JSON
	dataContentJSON, err := tool.StringifyJSON(encryptedData)
	if err != nil {
		return nil, err
	}

	// 先Base64编码，再使用商户私钥证书加密
	dataContentBase64 := base64.StdEncoding.EncodeToString([]byte(dataContentJSON))
	dataContentEncrypted, err := tool.RsaEncryptByPrivateKey(dataContentBase64, c.config.PrivateKey)
	if err != nil {
		return nil, err
	}

	// 构建请求参数
	reqParams := map[string]interface{}{
		"version":      "4.0.0.0",
		"member_id":    c.config.MemberId,
		"terminal_id":  c.config.TerminalId,
		"txn_type":     TXN_TYPE_REFUND,
		"txn_sub_type": TXN_SUB_TYPE_REFUND_QUERY,
		"data_type":    "json",
		"data_content": dataContentEncrypted,
	}

	// 发送请求
	responseStr, err := sendRequest(
		c.config.GetBaseURL(),
		reqParams,
	)
	if err != nil {
		return nil, err
	}

	// 解析并解密响应数据
	resp := &RefundQueryResponse{}
	if err := parseAndDecryptResponse(c.config, responseStr, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
