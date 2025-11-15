package bct3

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountBindRelationQueryRequest(t *testing.T) {
	memberId := "102005245"   //商户号
	terminalId := "200005972" //终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     //商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") //宝付公钥
	priKeyPass := "123456"

	configParams := &BCT3ConfigParams{
		MemberId:           memberId,
		TerminalId:         terminalId,
		IsProdMode:         false,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewBCT3Config(configParams)
	if err != nil {
		t.Errorf("NewBCT3Config failed: %v", err)
	}

	// 平台编号（客户账户号）
	contractNo := "CM610000000000174078" //"CM610000000000174898"

	req := &AccBindRelationQueryReq{
		RequestNo:  tool.GetRequestNo(),
		ContractNo: contractNo,
	}

	response := &AccBindRelationQueryResp{}

	response, err = AccountBindRelationQueryRequest(config, req)

	respStr, _ := tool.StringifyJSON(response)
	log.Printf("AccountBindRelationQueryRequest response: %s", respStr)

	if err != nil {
		t.Errorf("AccountBindRelationQueryRequest failed: %v", err)
	}
}
