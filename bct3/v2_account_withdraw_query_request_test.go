package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestAccountWithdrawQueryRequest(t *testing.T) {
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

	// 生成唯一流水号（查询时需要使用原交易的流水号）
	// 提示：这里应该使用之前调用取现接口时的 transSerialNo
	transSerialNo := "202604300028057882583867226" // 示例：使用之前取现交易的流水号

	// 转账结果查询（取现）请求参数
	req := &AccWithdrawQueryReq{
		TransSerialNo: transSerialNo,                   // 原商户流水号
		TradeTime:     time.Now().Format("2006-01-02"), // 交易时间 yyyy-MM-dd
	}

	// 执行取现查询请求
	resp, err := AccountWithdrawQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountWithdrawQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("取现查询响应: %+v\n", resp)

	// 检查查询结果
	t.Logf("取现查询请求成功")
	t.Logf("商户号: %s", resp.MemberId)
	t.Logf("原商户流水号: %s", resp.TransSerialNo)
	t.Logf("商户客户号: %s", resp.ContractNo)
	t.Logf("订单状态: %d", resp.State)
	t.Logf("转账金额: %.2f", resp.TransMoney)
	t.Logf("转账手续费: %.2f", resp.TransFee)
	t.Logf("转账交易时金额: %.2f", resp.TransferTotalAmount)

	// 根据状态判断
	switch resp.State {
	case WITHDRAW_QUERY_STATE_SUCCESS:
		t.Logf("取现成功")
		t.Logf("订单号: %d", resp.OrderId)
		t.Logf("成功时间: %s", resp.SuccessTime)
	case WITHDRAW_QUERY_STATE_PROCESSING:
		t.Logf("取现处理中，需要后续继续查询确认结果")
	case WITHDRAW_QUERY_STATE_FAILURE:
		t.Logf("取现失败: %s", resp.TransRemark)
	case WITHDRAW_QUERY_STATE_REFUNDED:
		t.Logf("取现已退回（可能因收款账户异常导致银行退票）")
		t.Logf("失败原因: %s", resp.TransRemark)
	}
}
