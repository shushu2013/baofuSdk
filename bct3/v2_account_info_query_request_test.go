package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestAccountInfoQueryRequest(t *testing.T) {
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

	// 账户信息查询请求参数 - 通过登录号查询
	req := &AccInfoQueryReq{
		Version: "1.0.0",
		LoginNo: "LN853033565871374", // 登录号(开户接口中的登录号)
		// ContractNo: "",                      // 客户账户号(可选)
		AccType: ACCOUNT_TYPE_ENTERPRISE, // 账户类型:1个人,2商户
		// PlatformNo: "102005245",             // 平台号(主商户号)
	}

	// 执行账户信息查询请求
	resp, err := AccountInfoQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountInfoQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户信息查询响应: %+v\n", resp)

	// 检查查询结果
	if resp.AccInfo.ContractNo == "" {
		t.Logf("账户信息查询成功，但未找到对应的账户记录")
	} else {
		t.Logf("账户信息查询成功，客户账户号: %s", resp.AccInfo.ContractNo)
		t.Logf("账户状态: %s", resp.AccInfo.AccState)
		t.Logf("客户名称: %s", resp.AccInfo.CustomerName)

		// 打印绑卡信息
		if len(resp.BindCardInfoList) > 0 {
			t.Logf("绑卡信息数量: %d", len(resp.BindCardInfoList))
			for i, bindCard := range resp.BindCardInfoList {
				t.Logf("第%d张卡 - 银行名称: %s, 持卡人: %s", i+1, bindCard.BankName, bindCard.CardUserName)
			}
		} else {
			t.Logf("未找到绑卡信息")
		}
	}
}
