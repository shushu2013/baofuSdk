package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestAccountBalanceQueryRequest(t *testing.T) {
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

	// 账户余额查询请求参数 - 查询二级户余额
	req := &AccBalanceQueryReq{
		Version:    "1.0.0",
		ContractNo: "CM610000000000174078",        // 替换为实际的二级客户号
		AccType:    ACCOUNT_BALANCE_TYPE_MERCHANT, // 2商户
	}

	// 执行余额查询请求
	resp, err := AccountBalanceQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountBalanceQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户余额查询响应: %+v\n", resp)

	// 检查查询结果
	t.Logf("账户余额查询成功")
	t.Logf("账簿可用余额: %.2f元", resp.AvailableBal)
	t.Logf("在途资金余额: %.2f元", resp.PendingBal)
	t.Logf("账簿余额: %.2f元", resp.CurrBal)

	// 余额说明
	t.Logf("说明: 账簿余额(%.2f) = 可用余额(%.2f) + 在途余额(%.2f) + 冻结金额",
		resp.CurrBal, resp.AvailableBal, resp.PendingBal)
}

func TestAccountBalanceQueryRequestForPlatform(t *testing.T) {
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

	// 账户余额查询请求参数 - 查询平台商户余额（一级户）
	req := &AccBalanceQueryReq{
		Version:    "1.0.0",
		ContractNo: "102005245",                   // 平台商户号（一级户）
		AccType:    ACCOUNT_BALANCE_TYPE_PLATFORM, // 4平台商户
	}

	// 执行余额查询请求
	resp, err := AccountBalanceQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountBalanceQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("平台商户余额查询响应: %+v\n", resp)

	// 检查查询结果
	t.Logf("平台商户余额查询成功")
	t.Logf("账簿可用余额: %.2f元", resp.AvailableBal)
	t.Logf("在途资金余额: %.2f元", resp.PendingBal)
	t.Logf("账簿余额: %.2f元", resp.CurrBal)
	t.Logf("说明: 一级户余额代表剩余可分账的总余额")
}
