package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementConfirmPay(t *testing.T) {
	memberId := "102004459"   //商户号
	terminalId := "100005196" //终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     //商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") //宝付公钥
	priKeyPass := "123456"

	configParams := &AgreementPayConfigParams{
		IsProdMode:         false,
		MemberId:           memberId,
		TerminalId:         terminalId,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewAgreementPayConfig(configParams)
	if err != nil {
		t.Errorf("NewAgreementPayConfig failed: %v", err)
	}

	smsCode := "123456" // 需要填入短信验证码
	// 预支付唯一码（从预支付接口返回）
	uniqueCode := "20241111111718313031100327173000" // 预支付一码
	// 短信验证码（用户收到短信后填写）

	// 信用卡信息（仅信用卡支付时需要，格式：信用卡有效期（yymm）|安全码）
	cardInfo := "" // 例如："2512|123" 表示有效期2025年12月，安全码123

	reqMap := &AgreementConfirmPayRequest{
		MsgId:      tool.GetMsgId(),
		UniqueCode: uniqueCode,
		SmsCode:    smsCode,
		CardInfo:   cardInfo,
	}

	resp, err := AgreementConfirmPay(config, reqMap)
	if err != nil {
		t.Errorf("AgreementConfirmPay failed: %v", err)
	}

	fmt.Println(resp)
}
