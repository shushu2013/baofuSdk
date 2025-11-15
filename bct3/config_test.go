package bct3

import (
	"log"
	"testing"
)

func TestBCT3Config_ParseCertSN(t *testing.T) {

	pfxPath := "/Users/wwb/项目文档/宝财通/BaoCaiTong3.0TestInfo_master1.0.0.1/BAOFU20240612_pri.pfx"     //商户私钥
	pubCerPath := "/Users/wwb/项目文档/宝财通/BaoCaiTong3.0TestInfo_master1.0.0.1/BAOFUP20240612_pub.cer" //宝付公钥
	priKeyPass := "123456"                                                                         //私钥密码

	params := &BCT3ConfigParams{
		MemberId:           "1234567890",
		TerminalId:         "1234567890",
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}

	config, err := NewBCT3Config(params)
	if err != nil {
		t.Errorf("NewBCT3Config() error = %v", err)
	}

	log.Printf("config: %v", config)
	log.Printf("SignSN: %s, NcrptnSN: %s", config.SignSN, config.NcrptnSN)
}
