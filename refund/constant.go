package refund

/**
 * 生产环境 - 统一接口地址
 */
const BASE_REFUND_API_URL = "https://public.baofoo.com/cutpayment/api/backTransRequest"

/**
 * 测试环境 - 统一接口地址
 */
const BASE_REFUND_API_TEST_URL = "https://vgw.baofoo.com/cutpayment/api/backTransRequest"

// 交易类型
const (
	TXN_TYPE_REFUND = "331" // 退款交易
)

// 交易子类
const (
	TXN_SUB_TYPE_REFUND       = "09" // 退款申请
	TXN_SUB_TYPE_REFUND_QUERY = "10" // 退款查询
)

// 退款类型
const (
	REFUND_TYPE_BAOFU_CASHIER = "1"  // 宝付收银台
	REFUND_TYPE_AUTH_PAY      = "2"  // 认证支付、代扣、快捷支付
	REFUND_TYPE_WECHAT        = "3"  // 微信支付
	REFUND_TYPE_ALIPAY        = "5"  // 支付宝支付
	REFUND_TYPE_AGREEMENT     = "8"  // 协议支付
	REFUND_TYPE_TRANSFER      = "18" // 转账支付
)

// 业务应答码
const BIZ_RESP_CODE_SUCCESS = "0000"

// 商户接口应答码
const (
	RESP_CODE_SUCCESS = "0000" // 成功
)
