package bct3

const Version = "0.0.1"

/**
 * 生产环境
 */
const BASE_BCT3_API_URL = "https://public.baofu.com/union-gw/napi"

/**
 * 测试环境
 */
const BASE_BCT3_API_TEST_URL = "https://vgw.baofoo.com/union-gw/napi"

// API 服务报文编号
const (
	// 用户开户
	SERVICE_ACCOUNT_OPEN = "BCT3-1002-001-01"
	// 开户结果查询
	SERVICE_ACCOUNT_OPEN_QUERY = "BCT3-1003-001-01"
	// 账户信息修改
	SERVICE_ACCOUNT_UPDATE = "BCT3-1005-001-01"
	// 账户信息查询
	SERVICE_ACCOUNT_QUERY = "BCT3-1008-001-01"
	// 账户升级
	SERVICE_ACCOUNT_UPGRADE = "T-1001-013-20"
	// 账户绑定关系操作
	SERVICE_ACCOUNT_BIND_RELATION_OPERATION = "T-1001-013-21"
	// 账户绑定关系查询
	SERVICE_ACCOUNT_BIND_RELATION_QUERY = "T-1001-013-22"
)

// 系统响应码
const (
	// 请求正常
	SYS_RESP_CODE_SUCCESS = "S_0000"
	// 请求受理失败
	SYS_RESP_CODE_REQUEST_FAILURE = "S_E_9001"
	// 请求受理结果未知
	SYS_RESP_CODE_REQUEST_UNKNOWN = "S_E_9002"
	// 商户信息不存在或状态不正常
	SYS_RESP_CODE_MERCHANT_NOT_FOUND = "S_E_0003"
	// 商户与终端号不匹配
	SYS_RESP_CODE_MERCHANT_TERMINAL_MISMATCH = "S_E_0004"
	// IP未绑定，请联系宝付
	SYS_RESP_CODE_IP_NOT_BOUND = "S_E_0005"
	// 明文参数格式或数据不正确
	SYS_RESP_CODE_PLAIN_PARAM_FORMAT_ERROR = "S_E_0001"
	// 明文参数解析失败
	SYS_RESP_CODE_PLAIN_PARAM_PARSE_ERROR = "S_E_0002"
	// 密文解密失败
	SYS_RESP_CODE_CIPHER_DECRYPT_ERROR = "S_E_0006"
	// 密文参数解析失败
	SYS_RESP_CODE_CIPHER_PARAM_PARSE_ERROR = "S_E_0007"
	// 接口服务报文不支持
	SYS_RESP_CODE_INTERFACE_NOT_SUPPORTED = "S_E_0010"
	// 验签失败
	SYS_RESP_CODE_SIGN_VERIFY_ERROR = "S_E_0011"
	// 解密失败
	SYS_RESP_CODE_DECRYPT_ERROR = "S_E_0012"
)

// retCode	含义	描述
// 0	失败	接口调用失败，异常或者参数校验失败。
// 1	成功	接口调用成功，具体业务是否成功。看具体的参数字段。
// 2	处理中	接口调用处理中，需要调用查询接口查询状态。
const (
	// 失败
	RET_CODE_FAILURE = 0
	// 成功
	RET_CODE_SUCCESS = 1
	// 处理中
	RET_CODE_PROCESSING = 2
)

// state	状态 1 成功 0 失败 -1 异常 2开户处理中
const (
	// 成功
	STATE_SUCCESS = 1
	// 失败
	STATE_FAILURE = 0
	// 异常
	STATE_EXCEPTION = -1
	// 开户处理中
	STATE_PROCESSING = 2
)

// 加密方式
const (
	// SM2withiSM3 国密
	VERIFY_TYPE_SM2 = "10"
	// RSA2048withSHA256 rsa加密
	VERIFY_TYPE_RSA = "11"
)

// 账户类型
const (
	// 个人
	ACCOUNT_TYPE_PERSONAL = 1
	// 企业/个体户
	ACCOUNT_TYPE_ENTERPRISE = 2
)

// 证件类型
const (
	// 身份证
	CERT_TYPE_ID_CARD = "ID"
	// 护照
	CERT_TYPE_PASSPORT = "PASSPORT"
	// 港澳通行证
	CERT_TYPE_HONG_KONG_AND_MACAO_PASS = "HONG_KONG_AND_MACAO_PASS"
	// 台湾同胞来往内地通行证
	CERT_TYPE_TAIWAN_TRAVEL_PERMIT = "TAIWAN_TRAVEL_PERMIT"
	// 营业执照
	CERT_TYPE_BUSINESS_LICENSE = "LICENSE"
)

// 上传的文件类型：
// 101 企业营业执照
// 102 银行开户许可证
// 104 法人身份证(正)
// 111 法人身份证(反)
const (
	// 企业营业执照
	FILE_TYPE_BUSINESS_LICENSE = "101"
	// 银行开户许可证
	FILE_TYPE_BANK_ACCOUNT_LICENSE = "102"
	// 法人身份证(正)
	FILE_TYPE_CORPORATE_ID_CARD_FRONT = "104"
	// 法人身份证(反)
	FILE_TYPE_CORPORATE_ID_CARD_BACK = "111"
)

// 提现卡修改
const OPT_TYPE_UPDATE_CARD = "02"

// 营业期限，长期99991231
const BUSINESS_EXECUTION_VALIDITY_PERIOD_LONG_TERM = "99991231"

// accState	String	2	M	用户状态：01开启、02关闭、03注销、04冻结
const (
	// 开启
	ACCOUNT_STATE_OPEN = "01"
	// 关闭
	ACCOUNT_STATE_CLOSED = "02"
	// 注销
	ACCOUNT_STATE_DELETED = "03"
	// 冻结
	ACCOUNT_STATE_FROZEN = "04"
)

// 账户绑定关系操作类型 01-新增、02-禁用
const (
	// 新增
	ACCOUNT_BIND_RELATION_OPERATION_TYPE_ADD = "01"
	// 禁用
	ACCOUNT_BIND_RELATION_OPERATION_TYPE_DISABLE = "02"
)

// 账户绑定关系状态 OPEN(“开启”),PENDING_OPEN(“待开启”),CLOSED(“关闭”);
const (
	// 开启
	ACCOUNT_BIND_RELATION_STATE_OPEN = "OPEN"
	// 待开启
	ACCOUNT_BIND_RELATION_STATE_PENDING_OPEN = "PENDING_OPEN"
	// 关闭
	ACCOUNT_BIND_RELATION_STATE_CLOSED = "CLOSED"
)
