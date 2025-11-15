package agreementPay

// 风控通用参数
type RiskBaseParam struct {
	// 行业类目，参见数据字典“行业类目”
	GoodsCategory string `json:"goodsCategory"`
	// 商户用户登录名，用户在商户系统中的登陆名（手机号、邮箱等标识）
	UserLoginId string `json:"userLoginId,omitempty"`
	// 用户邮箱，用户在商户系统中注册的邮箱
	UserEmail string `json:"userEmail,omitempty"`
	// 绑定手机号，商户系统中绑定手机号
	UserMobile string `json:"userMobile,omitempty"`
	// 用户注册姓名
	RegisterUserName string `json:"registerUserName,omitempty"`
	// 是否实名认证，1是 0不是
	IdentifyState string `json:"identifyState,omitempty"`
	// 用户身份证号
	UserIdNo string `json:"userIdNo,omitempty"`
	// 注册时间，格式：YYYYMMDDHHMMSS
	RegisterTime string `json:"registerTime,omitempty"`
	// 注册IP，用户在商户端注册时留存的IP
	RegisterIp string `json:"registerIp,omitempty"`
	// 持卡人姓名
	ChName string `json:"chName,omitempty"`
	// 持卡人身份证号
	ChIdNo string `json:"chIdNo,omitempty"`
	// 持卡人银行卡号
	ChCardNo string `json:"chCardNo,omitempty"`
	// 持卡人手机
	ChMobile string `json:"chMobile,omitempty"`
	// 持卡人支付IP，持卡人在支付时的IP地址；如无法获取有效IP，请传127.0.0.1
	ChPayIp string `json:"chPayIp"`
	// 设备指纹订单号，生成设备指纹的订单号(用于快捷)，如与支付订单号一致则传相同值
	DeviceOrderNo string `json:"deviceOrderNo,omitempty"`
}

// 电商风控参数
type EcommerceRiskParam struct {
	RiskBaseParam
	// 收货人姓名
	CsName string `json:"csName"`
	// 收货人手机，不包含电话国家码的本地号码
	CsMobile string `json:"csMobile"`
	// 收货人省份，参见我司省份代码
	CsProvince string `json:"csProvince"`
	// 收货人城市，参见我司城市代码
	CsCity string `json:"csCity"`
	// 收货人地址，不包含国家、省份/州、城市信息的详细地址
	CsAddress string `json:"csAddress"`
	// 商品名称，多种商品，半角分号分隔
	ProdNameList string `json:"prodNameList"`
	// 商品类目，多种商品规则同上
	ProdTypeList string `json:"prodTypeList"`
	// 商品数量，多种商品规则同上
	ProdQtyList string `json:"prodQtyList"`
	// 商品总价，数字(15位总长度，2位小数)，多种商品规则同上
	ProdAmtList string `json:"prodAmtList"`
	// 商品单价
	ProPrice string `json:"proPrice"`
	// 物流单号
	LogOrderNo string `json:"logOrderNo"`
	// 商户销售模式，自营/非自营
	SaleMode string `json:"saleMode"`
}

// 直接支付类交易
// https://docs.baofu.com/docs/interface_document/protocolPay-rsa
type AgreementDirectPayRequest struct {
	// 报文流水号，每次请求均不可重复
	MsgId string `json:"msg_id"`
	// 商户订单号，唯一订单号，8-50 位字母和数字,未支付成功的订单号可重复提交，重复提交时交易参数不得发生变化
	TransId string `json:"trans_id"`
	// 用户ID，用户在商户平台唯一ID
	UserId string `json:"user_id"`
	// 签约协议号
	ProtocolNo string `json:"protocol_no"`
	// 交易金额，单位：分
	TxnAmt string `json:"txn_amt"`
	// 卡信息
	CardInfo string `json:"card_info"`
	// 风控参数
	RiskItem string `json:"risk_item"`
	// 交易成功通知地址
	ReturnUrl string `json:"return_url"`
}

type AgreementPreBindBankCardRequest struct {
	// 报文流水号，每次请求均不可重复
	MsgId string `json:"msg_id"`
	// 用户ID，用户在商户平台唯一ID
	UserId string `json:"user_id"`
	// 卡类型
	CardType string `json:"card_type"`
	// 证件类型
	IdCardType string `json:"id_card_type"`
	// 卡信息，银行卡号|持卡人姓名|证件号|手机号|银行卡安全码|银行卡有效（yymm）,安全码，有效期非必填
	AccInfo string `json:"acc_info"`
}

type AgreementConfirmBindBankCardRequest struct {
	// 报文流水号，每次请求均不可重复
	MsgId string `json:"msg_id"`
	// 预签约唯一码
	UniqueCode string `json:"unique_code"`
	// 短信验证码
	SmsCode string `json:"sms_code"`
}

type AgreementQueryBindBankCardRequest struct {
	// 报文流水号，每次请求均不可重复
	MsgId string `json:"msg_id"`
	// 用户ID，用户在商户平台唯一ID
	UserId string `json:"user_id"`
	// 银行卡号，与user_id必须其中一个有值
	AccNo string `json:"acc_no"`
}

// 协议支付- 解除银行卡绑定请求参数
type AgreementUnBindBankCardRequest struct {
	// 报文流水号，每次请求均不可重复
	MsgId string `json:"msg_id"`
	// 用户ID，用户在商户平台唯一ID
	UserId string `json:"user_id"`
	// 签约协议号
	ProtocolNo string `json:"protocol_no"`
}
