package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestAccountOpenQueryRequest(t *testing.T) {
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

	// 开户查询请求参数
	req := &AccOpenQueryReq{
		// PlatformNo: "102005245",             // 平台商户号(主商户号)
		LoginNo: "LN853033565871374",     // 登录号(开户接口中的登录号)
		AccType: ACCOUNT_TYPE_ENTERPRISE, // 账户类型:1个人,2商户
	}

	// 执行开户查询请求
	resp, err := AccountOpenQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountOpenQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("开户查询响应: %+v\n", resp)

	// 检查响应码
	if resp.RetCode != RET_CODE_SUCCESS {
		t.Errorf("开户查询失败，错误码: %s, 错误信息: %s", resp.ErrorCode, resp.ErrorMsg)
		return
	}

	// 检查查询结果
	if resp.Result.ContractNo == "" {
		t.Logf("开户查询成功，但未找到对应的开户记录")
	} else {
		t.Logf("开户查询成功，客户账户号: %s", resp.Result.ContractNo)
	}
}

func TestAccountOpenQueryRequest_PersonalAccount(t *testing.T) {
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

	// 个人账户查询
	req := &AccOpenQueryReq{
		PlatformNo: "102005245",
		LoginNo:    "LNwQRM6iU22YYSgnd",   // 个人账户登录号
		AccType:    ACCOUNT_TYPE_PERSONAL, // 个人账户类型
	}

	// 执行开户查询请求
	resp, err := AccountOpenQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountOpenQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("个人账户开户查询响应: %+v\n", resp)

	// 检查响应码
	if resp.RetCode != RET_CODE_SUCCESS {
		t.Logf("个人账户开户查询失败，错误码: %s, 错误信息: %s", resp.ErrorCode, resp.ErrorMsg)
	} else {
		t.Logf("个人账户开户查询成功，客户账户号: %s", resp.Result.ContractNo)
	}
}
