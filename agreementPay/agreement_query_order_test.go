package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementQueryOrder(t *testing.T) {
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

	// 商户原始订单号（预支付或确认支付时提交的订单号）
	origTransId := "TIDB4D61A59701C039A8739"
	// 交易日期（订单提交的日期时间）
	origTradeDate := tool.FormatDateTime(time.Now(), true)

	reqMap := &AgreementQueryOrderRequest{
		MsgId:         tool.GetMsgId(),
		OrigTransId:   origTransId,
		OrigTradeDate: origTradeDate,
	}

	resp, err := AgreementQueryOrder(config, reqMap)
	if err != nil {
		t.Errorf("AgreementQueryOrder failed: %v", err)
	}

	fmt.Println(resp)
}
