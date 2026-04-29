package refund

// RefundRequest 退款请求参数
type RefundRequest struct {
	// 退款类型，参见常量：REFUND_TYPE_BAOFU_CASHIER(1:宝付收银台)、REFUND_TYPE_AUTH_PAY(2:认证支付、代扣、快捷支付)、REFUND_TYPE_WECHAT(3:微信支付)、REFUND_TYPE_ALIPAY(5:支付宝支付)、REFUND_TYPE_AGREEMENT(8:协议支付)、REFUND_TYPE_TRANSFER(18:转账支付)
	RefundType string `json:"refund_type"`
	// 原商户订单号
	TransId string `json:"trans_id"`
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 退款商户流水号
	TransSerialNo string `json:"trans_serial_no"`
	// 退款原因
	RefundReason string `json:"refund_reason"`
	// 退款金额，单位：分
	RefundAmt string `json:"refund_amt"`
	// 退款发起时间，14位定长，格式：年年年年月月日日时时分分秒秒
	RefundTime string `json:"refund_time"`
	// 附加字段，长度不超过128位
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 请求方保留域
	ReqReserved string `json:"req_reserved,omitempty"`
	// 服务器通知商户地址
	NoticeUrl string `json:"notice_url,omitempty"`
	// 营销组合支付分账退款信息
	UnionRefundInfo string `json:"union_refund_info,omitempty"`
	// 营销组合支付已分账退款信息
	PartShareRefundInfo string `json:"part_share_refund_info,omitempty"`
	// 营销组合支付分账退款是否垫资 传1：是，不传或传0：否
	ShareRefundAdvanceFlag string `json:"share_refund_advance_flag,omitempty"`
	// 分账退款信息
	ShareRefundInfo string `json:"share_refund_info,omitempty"`
}

// RefundEncryptedData 加密数据结构体
type RefundEncryptedData struct {
	// 终端号
	TerminalId string `json:"terminal_id"`
	// 商户号
	MemberId string `json:"member_id"`
	// 交易子类
	TxnSubType string `json:"txn_sub_type"`
	// 退款类型
	RefundType string `json:"refund_type"`
	// 原商户订单号
	TransId string `json:"trans_id"`
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 退款商户流水号
	TransSerialNo string `json:"trans_serial_no"`
	// 退款原因
	RefundReason string `json:"refund_reason"`
	// 退款金额，单位：分
	RefundAmt string `json:"refund_amt"`
	// 退款发起时间
	RefundTime string `json:"refund_time"`
	// 附加字段
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 请求方保留域
	ReqReserved string `json:"req_reserved,omitempty"`
	// 服务器通知商户地址
	NoticeUrl string `json:"notice_url,omitempty"`
	// 营销组合支付分账退款信息
	UnionRefundInfo string `json:"union_refund_info,omitempty"`
	// 营销组合支付已分账退款信息
	PartShareRefundInfo string `json:"part_share_refund_info,omitempty"`
	// 营销组合支付分账退款是否垫资
	ShareRefundAdvanceFlag string `json:"share_refund_advance_flag,omitempty"`
	// 分账退款信息
	ShareRefundInfo string `json:"share_refund_info,omitempty"`
}

// RefundResponse 退款响应参数
type RefundResponse struct {
	// 应答码
	RespCode string `json:"resp_code"`
	// 应答消息
	RespMsg string `json:"resp_msg"`
	// 退款宝付业务流水号
	RefundBusinessNo string `json:"refund_business_no"`
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 退款金额，单位：分
	RefundAmt string `json:"refund_amt"`
	// 终端号
	TerminalId string `json:"terminal_id"`
	// 商户号
	MemberId string `json:"member_id"`
	// 数据类型
	DataType string `json:"data_type"`
	// 交易类型
	TxnType string `json:"txn_type"`
	// 交易子类
	TxnSubType string `json:"txn_sub_type"`
	// 版本号
	Version string `json:"version,omitempty"`
	// 附加字段
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 预留字段
	ReqReserved string `json:"req_reserved,omitempty"`
}

// RefundQueryRequest 退款查询请求参数
type RefundQueryRequest struct {
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 商户流水号
	TransSerialNo string `json:"trans_serial_no"`
	// 附加字段，长度不超过128位
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 请求方保留域
	ReqReserved string `json:"req_reserved,omitempty"`
}

// RefundQueryEncryptedData 退款查询加密数据结构体
type RefundQueryEncryptedData struct {
	// 交易子类
	TxnSubType string `json:"txn_sub_type"`
	// 商户号
	MemberId string `json:"member_id"`
	// 终端号
	TerminalId string `json:"terminal_id"`
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 商户流水号
	TransSerialNo string `json:"trans_serial_no"`
	// 附加字段
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 请求方保留域
	ReqReserved string `json:"req_reserved,omitempty"`
}

// RefundQueryResponse 退款查询响应参数
type RefundQueryResponse struct {
	// 应答码
	RespCode string `json:"resp_code"`
	// 应答信息
	RespMsg string `json:"resp_msg"`
	// 退款商户订单号
	RefundOrderNo string `json:"refund_order_no"`
	// 成功退款金额，单位：分
	RefundAmt string `json:"refund_amt"`
	// 商户号
	MemberId string `json:"member_id"`
	// 终端号
	TerminalId string `json:"terminal_id"`
	// 交易类型
	TxnType string `json:"txn_type"`
	// 交易子类
	TxnSubType string `json:"txn_sub_type"`
	// 数据类型
	DataType string `json:"data_type"`
	// 版本号
	Version string `json:"version,omitempty"`
	// 附加字段
	AdditionalInfo string `json:"additional_info,omitempty"`
	// 预留字段
	ReqReserved string `json:"req_reserved,omitempty"`
	// 退回活动金额
	ActivityRefundAmt string `json:"activity_refund_amt,omitempty"`
}

// IsSuccess 判断退款是否成功（应答码为0000）
func (r *RefundResponse) IsSuccess() bool {
	return r.RespCode == RESP_CODE_SUCCESS
}

// IsProcessing 判断退款是否处理中或需要后续查询
func (r *RefundResponse) IsProcessing() bool {
	// 处理中或未知，需要后续查询的应答码
	processingCodes := map[string]bool{
		"BF00100": true, // 系统异常，请联系宝付
		"BF00112": true, // 系统繁忙，请稍后再试
		"BF00113": true, // 交易结果未知，请稍后查询
		"BF00115": true, // 交易处理中，请稍后查询
		"BF00202": true, // 交易超时，请稍后查询
		"BF00203": true, // 退款交易已受理
		"BF00244": true, // 退款订单创建失败
		"BF00307": true, // 退款处理中
		"BF00384": true, // 银行处理中
	}
	return processingCodes[r.RespCode]
}

// NeedQuery 判断是否需要发起查询（处理中或未知状态）
func (r *RefundResponse) NeedQuery() bool {
	return r.IsProcessing() || r.IsSuccess()
}

// IsSuccess 判断退款查询是否成功（应答码为0000）
func (r *RefundQueryResponse) IsSuccess() bool {
	return r.RespCode == RESP_CODE_SUCCESS
}

// IsProcessing 判断退款查询是否处理中或需要继续查询
func (r *RefundQueryResponse) IsProcessing() bool {
	// 处理中或未知，需要后续查询的应答码
	processingCodes := map[string]bool{
		"BF00100": true, // 系统异常，请联系宝付
		"BF00112": true, // 系统繁忙，请稍后再试
		"BF00113": true, // 交易结果未知，请稍后查询
		"BF00115": true, // 交易处理中，请稍后查询
		"BF00202": true, // 交易超时，请稍后查询
		"BF00203": true, // 退款交易已受理
		"BF00244": true, // 退款订单创建失败
		"BF00307": true, // 退款处理中
		"BF00384": true, // 银行处理中
	}
	return processingCodes[r.RespCode]
}

// NeedQuery 判断是否需要继续发起查询（处理中或未知状态）
func (r *RefundQueryResponse) NeedQuery() bool {
	return r.IsProcessing()
}
