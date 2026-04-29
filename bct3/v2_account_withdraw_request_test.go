package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountWithdrawRequest(t *testing.T) {
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

	// 生成唯一流水号
	transSerialNo := tool.GetTransSerialNo()

	// 转账（取现）请求参数
	req := &AccWithdrawReq{
		ContractNo:    "CM610000000000174078",                // 客户账户号
		TransSerialNo: transSerialNo,                         // 商户流水号
		DealAmount:    0.01,                                  // 提现金额，单位：元
		ReturnUrl:     "https://example.com/withdraw/notify", // 提现结果异步通知地址
		TransAbstract: "测试取现",                                // 摘要
	}

	// 执行取现请求
	resp, err := AccountWithdrawRequest(config, req)
	if err != nil {
		t.Errorf("AccountWithdrawRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("取现响应: %+v\n", resp)

	// 检查取现结果
	t.Logf("取现请求成功")
	t.Logf("请求流水号: %s", resp.TransSerialNo)
	t.Logf("客户账户号: %s", resp.ContractNo)
	t.Logf("订单状态: %d", resp.State)

	// 根据状态判断
	switch resp.State {
	case WITHDRAW_STATE_SUCCESS:
		t.Logf("取现受理成功")
	case WITHDRAW_STATE_FAILURE:
		t.Logf("取现受理失败: %s", resp.TransRemark)
	}
}
