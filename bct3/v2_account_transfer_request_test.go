package bct3

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestAccountTransferRequest(t *testing.T) {
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
	transSerialNo := fmt.Sprintf("TR%s%s", time.Now().Format("20060102150405"), "001")

	// 转账请求参数
	req := &AccTransferReq{
		Version:       "1.0.0",
		PayerNo:       "CM610000000000174078", // 替换为实际的付款方二级子商户号
		PayeeNo:       "CM610000000000174898", // 替换为实际的收款方二级子商户号
		TransSerialNo: transSerialNo,
		AccountType:   ACCOUNT_TYPE_BALANCE, // BALANCE-余额户 TRANSIT-在途户
		DealAmount:    "0.01",               // 转账金额，单位：元
		Remark:        "测试转账",               // 转账附言
	}

	// 执行转账请求
	resp, err := AccountTransferRequest(config, req)
	if err != nil {
		t.Errorf("AccountTransferRequest failed: %v", err)
		return
	}

	// 打印响应结果
	fmt.Printf("转账响应: %+v\n", resp)

	// 检查转账结果
	t.Logf("转账请求成功")
	t.Logf("请求流水号: %s", resp.TransSerialNo)
	t.Logf("业务流水号: %s", resp.BusinessNo)
	t.Logf("付款方: %s", resp.PayerNo)
	t.Logf("收款方: %s", resp.PayeeNo)
	t.Logf("转账金额: %.2f", resp.DealAmount)
	t.Logf("手续费: %.2f", resp.FeeAmount)
	t.Logf("订单状态: %d", resp.State)

	// 根据状态判断
	switch resp.State {
	case TRANSFER_STATE_SUCCESS:
		t.Logf("转账成功")
	case TRANSFER_STATE_PROCESSING:
		t.Logf("转账处理中，需要后续查询确认结果")
	case TRANSFER_STATE_FAILURE:
		t.Logf("转账失败: %s", resp.TransRemark)
	}
}
