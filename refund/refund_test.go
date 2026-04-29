package refund

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestRefund(t *testing.T) {
	memberId := "102004459"   // 商户号
	terminalId := "100005196" // 终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     // 商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") // 宝付公钥
	priKeyPass := "123456"

	configParams := &RefundConfigParams{
		IsProdMode:         false,
		MemberId:           memberId,
		TerminalId:         terminalId,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewRefundConfig(configParams)
	if err != nil {
		t.Errorf("NewRefundConfig failed: %v", err)
		return
	}

	// 创建退款客户端
	client := NewRefundClient(config)

	// 构建退款请求
	req := &RefundRequest{
		RefundType:    REFUND_TYPE_AGREEMENT,   // 8:协议支付
		TransId:       "23456788",              // 原商户发起的支付订单号，仅支持1年内校交易退款
		RefundOrderNo: "RON7012884951167302",   // 退款时商户端生成的订单号
		TransSerialNo: tool.GetTransSerialNo(), // 退款商户流水号
		RefundReason:  "用户退款",
		RefundAmt:     "1",                                 // 退款金额，单位：分
		RefundTime:    time.Now().Format("20060102150405"), // 退款发起时间
		NoticeUrl:     "",
	}

	resp, err := client.Refund(req)
	if err != nil {
		t.Errorf("Refund failed: %v", err)
		return
	}

	fmt.Printf("退款成功: %+v\n", resp)
}

func TestRefundQuery(t *testing.T) {
	memberId := "102004459"   // 商户号
	terminalId := "100005196" // 终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     // 商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") // 宝付公钥
	priKeyPass := "123456"

	configParams := &RefundConfigParams{
		IsProdMode:         false,
		MemberId:           memberId,
		TerminalId:         terminalId,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewRefundConfig(configParams)
	if err != nil {
		t.Errorf("NewRefundConfig failed: %v", err)
		return
	}

	// 创建退款客户端
	client := NewRefundClient(config)

	// 构建退款查询请求
	req := &RefundQueryRequest{
		RefundOrderNo: "REFUND20240101120000", // 退款商户订单号
		TransSerialNo: tool.GetMsgId(),        // 商户流水号
	}

	resp, err := client.RefundQuery(req)
	if err != nil {
		t.Errorf("RefundQuery failed: %v", err)
		return
	}

	fmt.Printf("退款查询成功: %+v\n", resp)
}

// TestRefundResponseAuxiliaryMethods 测试退款响应辅助方法
func TestRefundResponseAuxiliaryMethods(t *testing.T) {
	testCases := []struct {
		name          string
		respCode      string
		expectSuccess bool
		expectProcess bool
		expectQuery   bool
	}{
		// 成功状态
		{"成功-0000", "0000", true, false, true},
		// 处理中状态
		{"处理中-BF00100", "BF00100", false, true, true},
		{"处理中-BF00112", "BF00112", false, true, true},
		{"处理中-BF00113", "BF00113", false, true, true},
		{"处理中-BF00115", "BF00115", false, true, true},
		{"处理中-BF00202", "BF00202", false, true, true},
		{"处理中-BF00203", "BF00203", false, true, true},
		{"处理中-BF00244", "BF00244", false, true, true},
		{"处理中-BF00307", "BF00307", false, true, true},
		{"处理中-BF00384", "BF00384", false, true, true},
		// 其他失败状态
		{"失败-其他错误码", "BF99999", false, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &RefundResponse{
				RespCode: tc.respCode,
				RespMsg:  "测试消息",
			}

			if resp.IsSuccess() != tc.expectSuccess {
				t.Errorf("IsSuccess() = %v, want %v", resp.IsSuccess(), tc.expectSuccess)
			}

			if resp.IsProcessing() != tc.expectProcess {
				t.Errorf("IsProcessing() = %v, want %v", resp.IsProcessing(), tc.expectProcess)
			}

			if resp.NeedQuery() != tc.expectQuery {
				t.Errorf("NeedQuery() = %v, want %v", resp.NeedQuery(), tc.expectQuery)
			}
		})
	}
}

// TestRefundQueryResponseAuxiliaryMethods 测试退款查询响应辅助方法
func TestRefundQueryResponseAuxiliaryMethods(t *testing.T) {
	testCases := []struct {
		name          string
		respCode      string
		expectSuccess bool
		expectProcess bool
		expectQuery   bool
	}{
		// 成功状态
		{"成功-0000", "0000", true, false, false},
		// 处理中状态
		{"处理中-BF00100", "BF00100", false, true, true},
		{"处理中-BF00112", "BF00112", false, true, true},
		{"处理中-BF00113", "BF00113", false, true, true},
		{"处理中-BF00115", "BF00115", false, true, true},
		{"处理中-BF00202", "BF00202", false, true, true},
		{"处理中-BF00203", "BF00203", false, true, true},
		{"处理中-BF00244", "BF00244", false, true, true},
		{"处理中-BF00307", "BF00307", false, true, true},
		{"处理中-BF00384", "BF00384", false, true, true},
		// 其他失败状态
		{"失败-其他错误码", "BF99999", false, false, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := &RefundQueryResponse{
				RespCode: tc.respCode,
				RespMsg:  "测试消息",
			}

			if resp.IsSuccess() != tc.expectSuccess {
				t.Errorf("IsSuccess() = %v, want %v", resp.IsSuccess(), tc.expectSuccess)
			}

			if resp.IsProcessing() != tc.expectProcess {
				t.Errorf("IsProcessing() = %v, want %v", resp.IsProcessing(), tc.expectProcess)
			}

			if resp.NeedQuery() != tc.expectQuery {
				t.Errorf("NeedQuery() = %v, want %v", resp.NeedQuery(), tc.expectQuery)
			}
		})
	}
}
