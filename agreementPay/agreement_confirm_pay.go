package agreementPay

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 协议支付-确认支付接口
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa
func AgreementConfirmPay(config *AgreementPayConfig, req *AgreementConfirmPayRequest) (map[string]string, error) {
	// 交易类型
	txnType := TRANS_TYPE_CONFIRM_PAY

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 创建AES密钥
	aesKey := tool.CreateAeskey(16)
	dgtlEnvlp, err := tool.EncryptByPublicKey(tool.Base64Encode("01|"+aesKey), config.PublicKey)
	if err != nil {
		return nil, err
	}

	// 格式：预支付唯一码|短信验证码，先BASE64后进行AES加密
	uniqueCodeWithSms := req.UniqueCode + "|" + req.SmsCode
	uniqueCodeEncrypted, err := tool.AesEncrypt(tool.Base64Encode(uniqueCodeWithSms), aesKey)
	if err != nil {
		return nil, err
	}

	cardInfo := req.CardInfo
	// 如果有卡信息（信用卡支付时需要），进行加密
	if cardInfo != "" {
		cardInfo, err = tool.AesEncrypt(tool.Base64Encode(req.CardInfo), aesKey)
		if err != nil {
			return nil, err
		}
	}

	reqMap := map[string]string{
		"send_time":        timestamp,
		"msg_id":           req.MsgId,
		"version":          "4.0.0.0",
		"terminal_id":      config.TerminalId,
		"txn_type":         txnType,
		"member_id":        config.MemberId,
		"dgtl_envlp":       dgtlEnvlp,
		"unique_code":      uniqueCodeEncrypted,
		"card_info":        cardInfo,
		"req_reserved1":    req.ReqReserved1,
		"req_reserved2":    req.ReqReserved2,
		"additional_info1": req.AdditionalInfo1,
		"additional_info2": req.AdditionalInfo2,
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
