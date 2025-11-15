package agreementPay

import (
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

// 协议支付- 预绑卡
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa#17cryr
func AgreementPreBindBankCard(config *AgreementPayConfig, req *AgreementPreBindBankCardRequest) (map[string]string, error) {
	// 交易类型
	txnType := TRANS_TYPE_PRE_BIND_CARD

	// 时间戳
	timestamp := tool.FormatDateTime(time.Now(), true)

	// 创建AES密钥
	aesKey := tool.CreateAeskey(16)
	dgtlEnvlp, err := tool.EncryptByPublicKey(tool.Base64Encode("01|"+aesKey), config.PublicKey)
	if err != nil {
		return nil, err
	}

	// 先BASE64后进行AES加密
	accInfo, err := tool.AesEncrypt(tool.Base64Encode(req.AccInfo), aesKey)
	if err != nil {
		return nil, err
	}

	reqMap := map[string]string{
		"send_time":    timestamp,
		"msg_id":       req.MsgId,
		"version":      "4.0.0.0",
		"terminal_id":  config.TerminalId,
		"txn_type":     txnType,
		"member_id":    config.MemberId,
		"dgtl_envlp":   dgtlEnvlp,
		"user_id":      req.UserId,
		"card_type":    req.CardType,
		"id_card_type": req.IdCardType,
		"acc_info":     accInfo,
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

		// 解密账户信息中的敏感数据
		if resp["unique_code"] != "" {
			uniqueCode, _ := tool.AesDecrypt(resp["unique_code"], rAesKey)
			resp["unique_code"], _ = tool.Base64DecodeStr(uniqueCode)
		}
	}

	return resp, nil
}
