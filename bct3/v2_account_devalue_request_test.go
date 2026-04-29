package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountDevalueRequest(t *testing.T) {
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
	transOrderNo := fmt.Sprintf("TO%s%s", time.Now().Format("20060102150405"), "001")
	orderNo := fmt.Sprintf("ON%s%s", time.Now().Format("20060102150405"), "001")

	// 账户减值请求参数
	req := &AccDevalueReq{
		TransSerialNo: transSerialNo,
		PlatformNo:    config.MemberId,      // 平台商户号
		TransOrderNo:  transOrderNo,         // 商户订单号
		OrderNo:       orderNo,              // 商户原支付交易订单号
		AccountType:   ACCOUNT_TYPE_BALANCE, // BALANCE-余额户 TRANSIT-在途户
		AcctInfo: []*AccDevalueInfo{
			{
				ContractNo:  "CM610000000000174078", // 二级客户号
				TransAmount: "0.01",                 // 减值金额
				Remark:      "测试减值1",
			},
			{
				ContractNo:  "CM610000000000174898", // 二级客户号
				TransAmount: "0.01",                 // 减值金额
				Remark:      "测试减值2",
			},
		},
	}

	// 执行账户减值请求
	resp, err := AccountDevalueRequest(config, req)
	if err != nil {
		t.Errorf("AccountDevalueRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("账户减值响应: %+v\n", resp)

	// 检查减值结果
	t.Logf("账户减值请求成功")
	t.Logf("请求流水号: %s", resp.TransSerialNo)
	t.Logf("请求日期: %s", resp.ReqTime)
	t.Logf("商户订单号: %s", resp.OrderNo)
	t.Logf("减值总金额: %s", resp.DealAmount)
	t.Logf("账本类型: %s", resp.AccountType)
	t.Logf("手续费: %s", resp.TransFee)
	t.Logf("手续费承担方: %s", resp.FeeMemberId)
	t.Logf("订单状态: %d", resp.State)

	// 根据状态判断
	switch resp.State {
	case DEVALUE_STATE_SUCCESS:
		t.Logf("账户减值成功")
		t.Logf("减值成功时间: %s", resp.SuccessTime)
	case DEVALUE_STATE_PROCESSING:
		t.Logf("账户减值处理中，需要后续查询确认结果")
	case DEVALUE_STATE_FAILURE:
		t.Logf("账户减值失败: %s", resp.TransRemark)
	}
}
