package agreementPay

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 协议支付-预支付接口
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa
func AgreementPrePay(config *AgreementPayConfig, req *AgreementPrePayRequest) (map[string]string, error) {
	// 交易类型
	txnType := TRANS_TYPE_PRE_PAY

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

	reqMap := map[string]string{
		"send_time":          timestamp,
		"msg_id":             req.MsgId,
		"version":            "4.0.0.0",
		"terminal_id":        config.TerminalId,
		"txn_type":           txnType,
		"member_id":          config.MemberId,
		"trans_id":           req.TransId,
		"dgtl_envlp":         dgtlEnvlp,
		"user_id":            req.UserId,
		"protocol_no":        protocolNo,
		"txn_amt":            req.TxnAmt,
		"risk_item":          req.RiskItem,
		"return_url":         req.ReturnUrl,
		"req_reserved1":      req.ReqReserved1,
		"req_reserved2":      req.ReqReserved2,
		"additional_info1":   req.AdditionalInfo1,
		"additional_info2":   req.AdditionalInfo2,
		"fee_member_id":      req.FeeMemberId,
		"call_fee_member_id": req.CallFeeMemberId,
		"platform_no":        req.PlatformNo,
		"sub_merchant_no":    req.SubMerchantNo,
		"payment_type":       req.PaymentType,
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

	// 解密敏感数据
	if resp["resp_code"] == RESP_CODE_SUCCESS {
		// 解密响应头中的数字信封
		envelopeKey, err := tool.DecryptByPrivateKey(resp["dgtl_envlp"], config.PrivateKey)
		if err != nil {
			return nil, err
		}
		envelopeKey, err = tool.Base64DecodeStr(envelopeKey)
		if err != nil {
			return nil, err
		}

		// 从数字信封中解析出 AES 密钥
		rAesKey, err := getAesKey(envelopeKey)
		if err != nil {
			return nil, err
		}

		// 解密预支付唯一码
		if resp["unique_code"] != "" {
			uniqueCode, _ := tool.AesDecrypt(resp["unique_code"], rAesKey)
			resp["unique_code"], _ = tool.Base64DecodeStr(uniqueCode)
		}
	}

	return resp, nil
}
