package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountRechargeDevalueQueryRequest(t *testing.T) {
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
	// 查询之前的订单号（需要替换为实际存在的订单号）
	// 提示：可以先调用账户加值或减值接口创建订单，然后使用该订单号查询
	transOrderNo := "TON48CEB68757A9C7220157"

	// 账户加减值查询请求参数
	req := &AccRechargeDevalueQueryReq{
		TransSerialNo: transSerialNo,
		PlatformNo:    config.MemberId,    // 平台商户号
		TransOrderNo:  transOrderNo,       // 原商户订单号
		OrderType:     ORDER_TYPE_DEVALUE, // 01:加值；02:减值
	}

	// 执行账户加减值查询请求
	resp, err := AccountRechargeDevalueQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountRechargeDevalueQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户加减值查询响应: %+v\n", resp)

	// 检查查询结果
	t.Logf("账户加减值查询请求成功")
	t.Logf("请求流水号: %s", resp.TransSerialNo)
	t.Logf("请求日期: %s", resp.ReqTime)
	t.Logf("原商户订单号: %s", resp.TransOrderNo)
	t.Logf("业务类型: %s", resp.OrderType)
	t.Logf("成功金额: %s", resp.DealAmount)
	t.Logf("账本类型: %s", resp.AccountType)
	t.Logf("手续费: %s", resp.TransFee)
	t.Logf("手续费承担方: %s", resp.FeeMemberId)
	t.Logf("订单状态: %d", resp.State)

	// 根据状态判断
	switch resp.State {
	case RECHARGE_STATE_SUCCESS:
		t.Logf("订单处理成功")
		t.Logf("记账完成时间: %s", resp.FinishTime)
	case RECHARGE_STATE_PROCESSING:
		t.Logf("订单处理中，需要后续继续查询确认结果")
	case RECHARGE_STATE_FAILURE:
		t.Logf("订单失败: %s", resp.TransRemark)
	}
}
