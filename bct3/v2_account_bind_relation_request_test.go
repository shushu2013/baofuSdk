package bct3

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountBindRelationRequest(t *testing.T) {
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

	// 账户绑定关系请求参数 - 新增绑定关系
	req := &AccBindRelationReq{
		RequestNo:       tool.GetRequestNo(),
		ContractNo:      "CM610000000000174898",                   // 平台编号
		UpperContractNo: "CM610000000000174078",                   // 上级平台编号
		OperationType:   ACCOUNT_BIND_RELATION_OPERATION_TYPE_ADD, // 操作类型:01-新增
	}

	response := &AccBindRelationResp{}

	response, err = AccountBindRelationRequest(config, req)

	respStr, _ := tool.StringifyJSON(response)
	log.Printf("AccountBindRelationRequest response: %s", respStr)

	if err != nil {
		t.Errorf("AccountBindRelationRequest failed: %v", err)
	}
}

func TestAccountBindRelationRequest_Disable(t *testing.T) {
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

	// 账户绑定关系请求参数 - 禁用绑定关系
	req := &AccBindRelationReq{
		Version:         "4.1.0",
		RequestNo:       tool.GetRequestNo(),
		ContractNo:      "CM610000000000174898",                       // 平台编号
		UpperContractNo: "CM610000000000174078",                       // 上级平台编号
		OperationType:   ACCOUNT_BIND_RELATION_OPERATION_TYPE_DISABLE, // 操作类型:02-禁用
	}

	response := &AccBindRelationResp{}

	response, err = AccountBindRelationRequest(config, req)

	respStr, _ := tool.StringifyJSON(response)
	log.Printf("AccountBindRelationRequest response: %s", respStr)

	if err != nil {
		t.Errorf("AccountBindRelationRequest failed: %v", err)
	}
}
