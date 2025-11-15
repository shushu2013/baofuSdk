package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementPreBindBankCard(t *testing.T) {
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

	cardinfo := "6222021001015343573|王勇|310115200501018559|18616092709||" //账户信息[银行卡号|持卡人姓名|证件号|手机号|银行卡安全码|银行卡有效期]

	reqMap := &AgreementPreBindBankCardRequest{
		MsgId:      tool.GetMsgId(),
		UserId:     "",
		CardType:   BANK_CARD_TYPE_DEBIT,
		IdCardType: ID_CARD_TYPE_IDENTITY,
		AccInfo:    cardinfo,
	}

	resp, err := AgreementPreBindBankCard(config, reqMap)
	if err != nil {
		t.Errorf("AgreementPreBindBankCard failed: %v", err)
	}

	fmt.Println(resp)
}
