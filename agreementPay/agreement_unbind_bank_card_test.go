package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementUnBindBankCard(t *testing.T) {
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

	protocolNo := "1202508160015007180016781119" //签约协议号（确认绑卡返回）

	reqMap := &AgreementUnBindBankCardRequest{
		MsgId:      tool.GetMsgId(),
		ProtocolNo: protocolNo,
	}

	resp, err := AgreementUnBindBankCard(config, reqMap)
	if err != nil {
		t.Errorf("AgreementUnBindBankCard failed: %v", err)
	}

	fmt.Println(resp)
}
