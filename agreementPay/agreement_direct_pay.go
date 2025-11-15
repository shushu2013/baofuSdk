package agreementPay

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 协议支付-直接支付接口
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa#f3r4p7
func AgreementDirectPay(config *AgreementPayConfig, req *AgreementDirectPayRequest) (map[string]string, error) {
	// 交易类型
	txnType := TRANS_TYPE_DIRECT_PAY

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 创建AES密钥
	aesKey := tool.CreateAeskey(16)
	dgtlEnvlp, err := tool.EncryptByPublicKey(tool.Base64Encode("01|"+aesKey), config.PublicKey)
	if err != nil {
		return nil, err
	}

	// 先BASE64后进行AES加密
	protocolNo, err := tool.AesEncrypt(tool.Base64Encode(req.ProtocolNo), aesKey)
	if err != nil {
		return nil, err
	}

	// 先BASE64后进行AES加密
	cardInfo, err := tool.AesEncrypt(tool.Base64Encode(req.CardInfo), aesKey)
	if err != nil {
		return nil, err
	}

	reqMap := map[string]string{
		"send_time":   timestamp,
		"msg_id":      req.MsgId,
		"version":     "4.0.0.0",
		"terminal_id": config.TerminalId,
		"txn_type":    txnType,
		"member_id":   config.MemberId,
		"trans_id":    req.TransId,
		"dgtl_envlp":  dgtlEnvlp,
		"user_id":     req.UserId,
		"protocol_no": protocolNo,
		"txn_amt":     req.TxnAmt,
		"card_info":   cardInfo,
		"risk_item":   req.RiskItem,
		"return_url":  req.ReturnUrl,
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
