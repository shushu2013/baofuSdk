package bct3

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountUpgradeRequest(t *testing.T) {
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

	//客户账户号
	contractNo := "CM610000000000174078"
	// 上传接口返回的流水号
	qualificationTransSerialNo := "969F95E6A5FB96D71E8E"

	// 创建测试请求数据
	req := &AccUpgradeReq{
		RequestNo:                       fmt.Sprintf("UP%s", tool.RandomStr(15)),
		ContractNo:                      contractNo,
		QualificationTransSerialNo:      qualificationTransSerialNo,
		Province:                        "上海市",
		City:                            "上海市",
		District:                        "虹口区",
		DetailedAddress:                 "辰大道14号",
		BusinessAddress:                 "上海市上海市虹口区星辰大道14号 科学爱琴海14栋291室",
		RegisteredAddress:               "上海市上海市崇明区滨海大道104号 北京绿洲8栋2132室",
		LegalPersonIdValidityPeriod:     "20440506",
		BusinessScope:                   "物流",
		EstablishmentDate:               "20201023",
		RegisteredCapital:               "600",
		BusinessExecutionValidityPeriod: BUSINESS_EXECUTION_VALIDITY_PERIOD_LONG_TERM,
		ContactMobile:                   "13851124784",
	}

	response := &AccUpgradeResp{}

	response, err = AccountUpgradeRequest(config, req)

	respStr, _ := tool.StringifyJSON(response)
	log.Printf("AccountUpgradeRequest response: %s", respStr)

	if err != nil {
		t.Errorf("AccountUpgradeRequest failed: %v", err)
	}
}
