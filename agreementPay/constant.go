package agreementPay

/**
 * 生产环境
 */
const BASE_AGREEMENT_PAY_API_URL = "https://public.baofoo.com/cutpayment/protocol/backTransRequest"

/**
 * 测试环境
 */
const BASE_AGREEMENT_PAY_API_TEST_URL = "https://vgw.baofoo.com/cutpayment/protocol/backTransRequest"

// 交易类型枚举
// 交易类型	交易描述
// 01	协议支付预绑卡类交易
// 02	协议支付确认绑卡类交易
// 03	查询绑定关系类交易
// 04	协议支付解除绑卡类交易
// 05	协议支付预支付类交易
// 06	协议支付确认支付类交易
// 07	协议支付订单查询类交易
// 08	协议支付直接支付类
// 66	网关签约申请
// 67	网关签约预申请
// 68	网关确认签约
// 78	网关签约流水结果查询
const (
	TRANS_TYPE_PRE_BIND_CARD     = "01" // 协议支付预绑卡类交易
	TRANS_TYPE_CONFIRM_BIND_CARD = "02" // 协议支付确认绑卡类交易
	TRANS_TYPE_QUERY_BIND_REL    = "03" // 查询绑定关系类交易
	TRANS_TYPE_CANCEL_BIND_CARD  = "04" // 协议支付解除绑卡类交易
	TRANS_TYPE_PRE_PAY           = "05" // 协议支付预支付类交易
	TRANS_TYPE_CONFIRM_PAY       = "06" // 协议支付确认支付类交易
	TRANS_TYPE_QUERY_ORDER       = "07" // 协议支付订单查询类交易
	TRANS_TYPE_DIRECT_PAY        = "08" // 协议支付直接支付类
	TRANS_TYPE_APPLY_SIGN        = "66" // 网关签约申请
	TRANS_TYPE_PRE_APPLY_SIGN    = "67" // 网关签约预申请
	TRANS_TYPE_CONFIRM_SIGN      = "68" // 网关确认签约
	TRANS_TYPE_QUERY_SIGN_FLOW   = "78" // 网关签约流水结果查询
)

// 卡类型
const (
	BANK_CARD_TYPE_DEBIT  = "101" // 借记卡
	BANK_CARD_TYPE_CREDIT = "102" // 信用卡
)

// 证件类型
const (
	ID_CARD_TYPE_IDENTITY = "01" // 身份证
	ID_CARD_TYPE_HONGKONG = "12" // 港澳居民居住证
	ID_CARD_TYPE_TAIWAN   = "13" // 台湾居民居住证
	ID_CARD_TYPE_FOREIGN  = "09" // 外国人永久居住证
)

// 行业类目（goodsCategory）
const (
	GOODS_CATEGORY_ECOMMERCE      = "01" // 电商
	GOODS_CATEGORY_FINANCE        = "02" // 互金消金
	GOODS_CATEGORY_TRAVEL         = "03" // 航旅
	GOODS_CATEGORY_HOTEL          = "04" // 酒店
	GOODS_CATEGORY_INSURANCE      = "05" // 保险
	GOODS_CATEGORY_GAME           = "06" // 游戏
	GOODS_CATEGORY_BULK_COMMODITY = "07" // 大宗
)

// 协议支付相关biz_resp_code业务应答码
const BIZ_RESP_CODE_SUCCESS = "0000"

// 商户接口应答码
// 应答码	resp_code
const (
	RESP_CODE_SUCCESS    = "S"  // 成功
	RESP_CODE_FAIL       = "F"  // 失败
	RESP_CODE_PROCESS    = "I"  // 处理中
	RESP_CODE_FAIL_QUERY = "FF" // 失败（支付结果查询类交易才会返回，表示订单查询参数错误或其他原因导致的订单查询失败，而非订单交易失败）
)
