package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementConfirmBindBankCard(t *testing.T) {
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

	uniqueCode := "202511150646163041881" //预签约唯一码
	// 短信验证码，测试环境随机6位数;生产环境验证码预绑卡成功后发到用户手机。确认绑卡时回传。
	smsCode := "123456" //短信验证码

	reqMap := &AgreementConfirmBindBankCardRequest{
		MsgId:      tool.GetMsgId(),
		UniqueCode: uniqueCode,
		SmsCode:    smsCode,
	}

	resp, err := AgreementConfirmBindBankCard(config, reqMap)
	if err != nil {
		t.Errorf("AgreementConfirmBindBankCard failed: %v", err)
	}

	fmt.Println(resp)
}
