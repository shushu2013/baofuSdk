package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountInfoUpdateRequest(t *testing.T) {
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

	contractNo := "CM610000000000174078" //客户账户号
	newCardNo := "96660136380784"        //账号

	//对私
	// mobileNo := "17762651048" //对私手机号

	//对公
	BankName := "工商银行"          //银行名称
	DepositBankProvince := "上海" //开户行省份
	DepositBankCity := "上海"     //开户行城市
	DepositBankName := "张江支行"   //开户支行名称

	// 账户信息修改请求参数 - 修改提现卡
	req := &AccInfoUpdateReq{
		AccType: ACCOUNT_TYPE_ENTERPRISE, // 账户类型:1个人,2商户
		OptType: OPT_TYPE_UPDATE_CARD,    // 操作类型:02:提现卡修改
		AccInfo: AccInfoUpdateInfo{
			ContractNo:          contractNo,              // 客户账户号
			TransSerialNo:       tool.GetTransSerialNo(), // 请求流水号
			CardNo:              newCardNo,               // 卡号 (修改卡号时必传)
			BankName:            BankName,                // 银行名称 (修改卡号时必传)
			DepositBankProvince: DepositBankProvince,     // 开户行省份 (修改卡号时必传)
			DepositBankCity:     DepositBankCity,         // 开户行城市 (修改卡号时必传)
			DepositBankName:     DepositBankName,         // 开户支行 (修改卡号时必传)
			// MobileNo:            mobileNo,            // 银行预留手机号，提现卡修改必填
		},
	}

	// 执行账户信息修改请求
	resp, err := AccountInfoUpdateRequest(config, req)
	if err != nil {
		t.Errorf("AccountInfoUpdateRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户信息修改响应: %+v\n", resp)

	t.Logf("账户信息修改成功，客户账户号: %s", resp.ContractNo)
}

func TestAccountInfoUpdateRequest_ContactInfo(t *testing.T) {
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

	contractNo := "CM610000000000174078" //客户账户号
	newCardNo := "96660136380784"        //账号

	//联系人手机号
	contactMobile := "17762651048" //联系人手机号

	//对公
	BankName := "工商银行"          //银行名称
	DepositBankProvince := "上海" //开户行省份
	DepositBankCity := "上海"     //开户行城市
	DepositBankName := "张江支行"   //开户支行名称

	// 账户信息修改请求参数 - 修改提现卡
	req := &AccInfoUpdateReq{
		AccType: ACCOUNT_TYPE_ENTERPRISE, // 账户类型:1个人,2商户
		OptType: "02",                    // 操作类型:02:提现卡修改
		AccInfo: AccInfoUpdateInfo{
			ContractNo:          contractNo,              // 客户账户号
			TransSerialNo:       tool.GetTransSerialNo(), // 请求流水号
			CardNo:              newCardNo,               // 卡号 (修改卡号时必传)
			BankName:            BankName,                // 银行名称 (修改卡号时必传)
			DepositBankProvince: DepositBankProvince,     // 开户行省份 (修改卡号时必传)
			DepositBankCity:     DepositBankCity,         // 开户行城市 (修改卡号时必传)
			DepositBankName:     DepositBankName,         // 开户支行 (修改卡号时必传)
			ContactMobile:       contactMobile,           // 联系人手机号【敏感信息（2.0入口使用）】
		},
	}

	// 执行账户信息修改请求
	resp, err := AccountInfoUpdateRequest(config, req)
	if err != nil {
		t.Errorf("AccountInfoUpdateRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户信息修改响应: %+v\n", resp)

	t.Logf("账户信息修改成功，客户账户号: %s", resp.ContractNo)
}
