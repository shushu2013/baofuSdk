package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestAccountTransferQueryRequest(t *testing.T) {
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
	// 提示：这里应该使用之前调用转账接口时的 transSerialNo
	transSerialNo := "TSNC393A74D6D8643F9864B" // 示例：使用之前转账交易的流水号

	// 转账结果查询（账户间）请求参数
	req := &AccTransferQueryReq{
		TransSerialNo: transSerialNo, // 原商户流水号
		TradeTime:     "2025-04-05",  // 交易时间 yyyy-MM-dd
	}

	// 执行转账查询请求
	resp, err := AccountTransferQueryRequest(config, req)
	if err != nil {
		t.Errorf("AccountTransferQueryRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("转账查询响应: %+v\n", resp)

	// 检查查询结果
	t.Logf("转账查询请求成功")
	t.Logf("请求流水号: %s", resp.TransSerialNo)
	t.Logf("业务流水号: %s", resp.BusinessNo)
	t.Logf("付款方: %s", resp.PayerNo)
	t.Logf("收款方: %s", resp.PayeeNo)
	t.Logf("转账金额: %.2f", resp.DealAmount)
	t.Logf("手续费: %.2f", resp.FeeAmount)
	t.Logf("订单状态: %d", resp.State)
	t.Logf("账户类型: %s", resp.AccountType)

	// 根据状态判断
	switch resp.State {
	case TRANSFER_QUERY_STATE_SUCCESS:
		t.Logf("转账成功")
		t.Logf("成功时间: %s", resp.SuccessTime)
	case TRANSFER_QUERY_STATE_PROCESSING:
		t.Logf("转账处理中，需要后续继续查询确认结果")
	case TRANSFER_QUERY_STATE_FAILURE:
		t.Logf("转账失败: %s", resp.TransRemark)
	}
}
