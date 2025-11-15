package bct3

// RequestHeader 请求头
// 对应宝付API文档中的header部分
// https://docs.baofu.com/docs/bct3/bct3Entrance
// https://docs.baofu.com/docs/bct3/bct3-1g42bpgpun6ss
type RequestHeader struct {
	// 商户号
	MemberID string `json:"memberId"`
	// 终端号
	TerminalID string `json:"terminalId"`
	// 报文发送时间，格式：yyyy-MM-dd HH:mm:ss
	Timestamp string `json:"timestamp"`
	// 加密方式：10–国密 11-RSA
	VerifyType string `json:"verifyType"`
	// 字符集：固定字符集UTF-8
	Charset string `json:"charset"`
	// 接口版本号 1.0
	Version string `json:"version"`
	// 签名证书序列号
	SignSN string `json:"signSN"`
	// 加密证书序列号
	NcrptnSN string `json:"ncrptnSN"`
	// 数字信封
	DgtlEnvlp string `json:"dgtlEnvlp,omitempty"`
}

// 请求参数
type RequestData struct {
	// 请求公共参数
	Header string `json:"header"`
	// 业务参数
	Body string `json:"body"`
	// 签名串
	Sign string `json:"sign"`
}

// ResponseHeader 响应头
type ResponseHeader struct {
	// 商户号
	MemberID string `json:"memberId"`
	// 终端号
	TerminalID string `json:"terminalId"`
	// 服务编号
	ServiceTp string `json:"serviceTp"`
	// 加密方式
	VerifyType string `json:"verifyType"`
	// 数字信封
	DgtlEnvlp string `json:"dgtlEnvlp"`
	// 签名证书序列号
	SignSN string `json:"signSN"`
	// 加密证书序列号
	NcrptnSN string `json:"ncrptnSN"`
	// 系统返回码
	SysRespCode string `json:"sysRespCode"`
	// 系统返回说明
	SysRespDesc string `json:"sysRespDesc"`
}

// 响应参数
type ResponseData struct {
	// 返回公共参数
	Header string `json:"header"`
	// 返回业务数据
	Body string `json:"body"`
	// 签名串
	Sign string `json:"sign"`
}

// 通知数据
type NotifyData struct {
	// 商户号
	MemberID string `json:"memberId"`
	// 终端号
	TerminalID string `json:"terminalId"`
	// JSON,data_content的明文组装格式
	DataType string `json:"dataType"`
	// 业务报文编号，例如：BCT3-1104-001-01
	ServiceTp string `json:"serviceTp"`
	// 数字信封
	DgtlEnvlp string `json:"dgtlEnvlp"`
	// 签名串
	Signature string `json:"signature"`
	// 业务数据
	DataContent string `json:"dataContent"`
}

// 开户用户信息
type AccOpenInfo struct {
	// 流水号每次请求不重复
	TransSerialNo string `json:"transSerialNo"`
	// 登录号商户端用户唯一标识，长度11位以上，商户自定义（不重复）
	LoginNo string `json:"loginNo"`
	// 邮箱
	Email string `json:"email"`
	// 是否个体户（ 企业为false，个体户为true）
	SelfEmployed bool `json:"selfEmployed"`
	// 个人：客户名称与持卡人姓名一致 企业：商户名称（营业执照上的名称）
	CustomerName string `json:"customerName"`
	// 商户名称别名
	AliasName string `json:"aliasName,omitempty"`
	// 个人：身份证号码，企业：营业执照证件号码
	CertificateNo string `json:"certificateNo"`
	// 证件类型
	CertificateType string `json:"certificateType"`
	// 法人姓名
	CorporateName string `json:"corporateName,omitempty"`
	// 法人证件类型
	CorporateCertType string `json:"corporateCertType,omitempty"`
	// 法人证件号码
	CorporateCertId string `json:"corporateCertId,omitempty"`
	// 法人手机号，当开个体户(selfEmployed=true)且绑定对私卡时必传【敏感信息（2.0入口使用）】
	CorporateMobile string `json:"corporateMobile,omitempty"`
	// 公司所属行业 见附录
	IndustryId string `json:"industryId,omitempty"`
	// 联系人姓名
	ContactName string `json:"contactName,omitempty"`
	// 联系人手机号【敏感信息（2.0入口使用）】
	ContactMobile string `json:"contactMobile,omitempty"`
	// 个人：对私卡号，企业：对公卡号，（个体可绑法人对私卡号）
	CardNo string `json:"cardNo"`
	// 银行名称
	BankName string `json:"bankName,omitempty"`
	// 开户行省份
	DepositBankProvince string `json:"depositBankProvince,omitempty"`
	// 开户行城市
	DepositBankCity string `json:"depositBankCity,omitempty"`
	// 开户行名称
	DepositBankName string `json:"depositBankName,omitempty"`
	// 注册资本
	RegisterCapital string `json:"registerCapital,omitempty"`
	// 持卡人姓名, 当开个体户且绑定对私卡时需传此字段,否则默认绑定对公卡
	CardUserName string `json:"cardUserName,omitempty"`
	// 平台号
	PlatformNo string `json:"platformNo,omitempty"`
	// 平台终端号
	PlatformTerminalId string `json:"platformTerminalId,omitempty"`
	// 资质文件流水,businessType为宝财通3非必填
	QualificationTransSerialNo string `json:"qualificationTransSerialNo,omitempty"`
	// 银行预留手机号
	MobileNo string `json:"mobileNo,omitempty"`
	// 是否需要上传附件,true/false
	NeedUploadFile bool `json:"needUploadFile"`
}

// 开户请求参数
type AccOpenReq struct {
	// 版本号 1.0.0
	Version string `json:"version"`
	// 账户类型:1-个人,2-企业/个体
	AccType int `json:"accType"`
	// 账户信息实体
	AccInfo AccOpenInfo `json:"accInfo"`
	// 通知地址
	NoticeUrl string `json:"noticeUrl"`
	// 宝财通3: BCT3
	BusinessType string `json:"businessType"`
}

// 账户基础响应参数
type AccBaseResp struct {
	// 版本号
	Version string `json:"version,omitempty"`
	// 返回码
	RetCode int `json:"retCode"`
	// 错误码
	ErrorCode string `json:"errorCode,omitempty"`
	// 错误信息
	ErrorMsg string `json:"errorMsg,omitempty"`
	// 备用字段1
	Back1 string `json:"back1,omitempty"`
	// 备用字段2
	Back2 string `json:"back2,omitempty"`
	// 备用字段3
	Back3 string `json:"back3,omitempty"`
}

// 开户响应数据
type AccOpenResp struct {
	AccBaseResp
	// 返回数据列表
	Result []*AccOpenRespResult `json:"result"`
}

// 开户响应 result 数据
type AccOpenRespResult struct {
	// 状态 1 成功 0 失败 -1 异常 2开户处理中
	State int `json:"state"`
	// 错误码
	ErrorCode string `json:"errorCode"`
	// 错误原因
	ErrorMsg string `json:"errorMsg"`
	// 请求流水号
	TransSerialNo string `json:"transSerialNo"`
	// 登录号
	LoginNo string `json:"loginNo"`
	// 商户名称
	CustomerName string `json:"customerName"`
	// 商户客户号
	ContractNo string `json:"contractNo"`
}

// 开户查询请求参数
type AccOpenQueryReq struct {
	// 版本号 1.0.0
	Version string `json:"version"`
	// 平台商户号(主商户号)
	PlatformNo string `json:"platformNo,omitempty"`
	// 登录号(开户接口中的登录号)
	LoginNo string `json:"loginNo"`
	// 账户类型:1个人,2商户
	AccType int `json:"accType"`
}

// 开户查询响应数据
type AccOpenQueryResp struct {
	AccBaseResp
	// 返回数据列表
	Result AccOpenQueryRespResult `json:"result"`
}

// 开户查询响应 result 数据
type AccOpenQueryRespResult struct {
	// 客户账户号
	ContractNo string `json:"contractNo"`
	// 客户账户名
	ContractName string `json:"contractName,omitempty"`
	// 客户名
	CustomerName string `json:"customerName,omitempty"`
	// 客户号
	CustomerNo string `json:"customerNo,omitempty"`
	// 客户类型
	CustomerType string `json:"customerType,omitempty"`
	// 证件类型 只能取”ID”或”LICENSE”
	CertificateType string `json:"certificateType,omitempty"`
	// 证件号码
	CertificateNo string `json:"certificateNo,omitempty"`
	// 平台号
	PlatformNo string `json:"platformNo,omitempty"`
	// 绑定手机号
	BindMobile string `json:"bindMobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
}

// 开户结果通知数据
type AccOpenNotifyResp struct {
	// 版本号
	Version string `json:"version"`
	// 类型:1-个人,2-企业,3-个体工商户
	MemberType int `json:"memberType"`
	// 请求流水号
	TransSerialNo string `json:"transSerialNo"`
	// 状态 1 成功 0 失败 -1 异常 2开户处理中
	State string `json:"state"`
	// 错误码
	ErrorCode string `json:"errorCode"`
	// 错误信息
	ErrorMsg string `json:"errorMsg"`
	// 登录号
	LoginNo string `json:"loginNo"`
	// 商户名称
	CustomerName string `json:"customerName"`
	// 商户客户号
	ContractNo string `json:"contractNo"`
}

// 账户信息查询请求参数
type AccInfoQueryReq struct {
	// 版本号 1.0.0
	Version string `json:"version"`
	// 登录号(无商户客户号必填)
	LoginNo string `json:"loginNo,omitempty"`
	// 客户账户号
	ContractNo string `json:"contractNo,omitempty"`
	// 账户类型:1个人,2商户
	AccType int `json:"accType"`
	// 平台号(主商户号)(无商户客户号必填)
	PlatformNo string `json:"platformNo,omitempty"`
}

// 账户信息查询响应数据
type AccInfoQueryResp struct {
	AccBaseResp
	// 账户信息，当retCode=1时有值
	AccInfo AccInfoQueryRespAccInfo `json:"accInfo,omitempty"`
	// 绑卡信息,当retCode=1是有值
	BindCardInfoList []*AccInfoQueryBindCardInfo `json:"bindCardInfoList,omitempty"`
}

// 账户信息查询响应 accInfo 数据
type AccInfoQueryRespAccInfo struct {
	// 客户账户号
	ContractNo string `json:"contractNo"`
	// 客户账户名
	ContractName string `json:"contractName,omitempty"`
	// 客户名
	CustomerName string `json:"customerName,omitempty"`
	// 客户号
	CustomerNo string `json:"customerNo,omitempty"`
	// 客户类型
	CustomerType string `json:"customerType,omitempty"`
	// 证件类型 只能取”ID”或”LICENSE”
	CertificateType string `json:"certificateType,omitempty"`
	// 证件号码（社会信用代码）
	CertificateNo string `json:"certificateNo,omitempty"`
	// 平台号
	PlatformNo string `json:"platformNo,omitempty"`
	// 邮箱，accType = 2（企业）时才会返回邮箱
	Email string `json:"email,omitempty"`
	// 账户状态 01开启、02关闭、03注销、04冻结
	AccState string `json:"accState"`
}

// 账户信息查询响应 bindCardInfoList 列表项数据
type AccInfoQueryBindCardInfo struct {
	// 客户名称
	CardUserName string `json:"cardUserName"`
	// 卡号【敏感信息（2.0入口使用）】
	CardNo string `json:"cardNo"`
	// 银行名称
	BankName string `json:"bankName"`
	// 预留手机号【敏感信息（2.0入口使用）】,accType=1 有值
	MobileNo string `json:"mobileNo"`
	// 开户行省份
	DepositBankProvince string `json:"depositBankProvince"`
	// 开户行城市
	DepositBankCity string `json:"depositBankCity"`
	// 开户支行名称
	DepositBankName string `json:"depositBankName"`
}

// 账户信息修改请求参数
type AccInfoUpdateReq struct {
	// 版本号 1.0.0
	Version string `json:"version"`
	// 账户类型:1个人,2商户
	AccType int `json:"accType"`
	// 操作类型:02:提现卡修改
	OptType string `json:"optType"`
	// 账户信息实体
	AccInfo AccInfoUpdateInfo `json:"accInfo"`
}

// 账户信息修改信息实体
type AccInfoUpdateInfo struct {
	// 客户账户号
	ContractNo string `json:"contractNo"`
	// 请求流水号
	TransSerialNo string `json:"transSerialNo"`
	// 卡号 (修改卡号时必传)【敏感信息（2.0入口使用）】
	CardNo string `json:"cardNo"`
	// 银行名称 (修改卡号时必传)
	BankName string `json:"bankName,omitempty"`
	// 开户行省份 (修改卡号时必传)
	DepositBankProvince string `json:"depositBankProvince,omitempty"`
	// 开户行城市 (修改卡号时必传)
	DepositBankCity string `json:"depositBankCity,omitempty"`
	// 开户支行 (修改卡号时必传)
	DepositBankName string `json:"depositBankName,omitempty"`
	// 联系人姓名
	ContactName string `json:"contactName,omitempty"`
	// 联系人手机号【敏感信息（2.0入口使用）】
	ContactMobile string `json:"contactMobile,omitempty"`
	// 法人手机号，当开个体户且绑定对私卡时必传【敏感信息（2.0入口使用）】
	CorporateMobile string `json:"corporateMobile,omitempty"`
	// 公司名称
	CustomerName string `json:"customerName,omitempty"`
	// 法人姓名
	CorporateName string `json:"corporateName,omitempty"`
	// 法人身份证号
	CorporateCertId string `json:"corporateCertId,omitempty"`
	// 银行预留手机号，提现卡修改必填【敏感信息（2.0入口使用）】
	MobileNo string `json:"mobileNo,omitempty"`
	// 持卡人姓名, 当开个体户且绑定对私卡时需传此字段,否则默认绑定对公卡
	CardUserName string `json:"cardUserName,omitempty"`
}

// 账户信息修改响应数据
type AccInfoUpdateResp struct {
	AccBaseResp
	// 商户客户号
	ContractNo string `json:"contractNo"`
}

// 账户升级请求参数
type AccUpgradeReq struct {
	// 版本号 4.1.0
	Version string `json:"version"`
	// 请求流水号
	RequestNo string `json:"requestNo"`
	// 平台号
	ContractNo string `json:"contractNo"`
	// 资质文件请求流水号，为文件上传接口返回的请求流水号
	QualificationTransSerialNo string `json:"qualificationTransSerialNo"`
	// 公司地址-省份
	Province string `json:"province"`
	// 公司地址-城市
	City string `json:"city"`
	// 公司地址-区
	District string `json:"district"`
	// 公司地址-详细地址
	DetailedAddress string `json:"detailedAddress"`
	// 经营地址
	BusinessAddress string `json:"businessAddress"`
	// 注册地址
	RegisteredAddress string `json:"registeredAddress"`
	// 法人证件有效期 yyyyMMdd
	LegalPersonIdValidityPeriod string `json:"legalPersonIdValidityPeriod"`
	// 企业经营范围
	BusinessScope string `json:"businessScope"`
	// 成立时间 yyyyMMdd
	EstablishmentDate string `json:"establishmentDate"`
	// 注册资本 单位万元
	RegisteredCapital string `json:"registeredCapital"`
	// 营业期限 yyyyMMdd,长期99991231
	BusinessExecutionValidityPeriod string `json:"businessExecutionValidityPeriod"`
	// 联系人手机号,用于接收合作协议电子签验证码。开户时未传此字段必填【敏感信息（2.0入口使用）】
	ContactMobile string `json:"contactMobile,omitempty"`
}

// 账户升级响应数据
type AccUpgradeResp struct {
	AccBaseResp
}

// 账户绑定关系请求参数
type AccBindRelationReq struct {
	// 版本号 4.1.0
	Version string `json:"version"`
	// 请求流水号
	RequestNo string `json:"requestNo"`
	// 平台编号
	ContractNo string `json:"contractNo"`
	// 上级平台编号
	UpperContractNo string `json:"upperContractNo"`
	// 01-新增、02-禁用
	OperationType string `json:"operationType"`
}

// 账户绑定关系响应数据
type AccBindRelationResp struct {
	AccBaseResp
}

// 账户绑定关系查询请求参数
type AccBindRelationQueryReq struct {
	// 版本号 4.1.0
	Version string `json:"version"`
	// 请求流水号
	RequestNo string `json:"requestNo"`
	// 平台编号
	ContractNo string `json:"contractNo"`
}

// 账户绑定关系查询响应数据
type AccBindRelationQueryResp struct {
	AccBaseResp
	// 账号绑定关系
	List []*AccBindRelationInfo `json:"list"`
}

// 账号绑定关系信息
type AccBindRelationInfo struct {
	// 平台编号
	ContractNo string `json:"contractNo"`
	// 上级平台编号
	UpperContractNo string `json:"upperContractNo"`
	// 状态 OPEN(“开启”),PENDING_OPEN(“待开启”),CLOSED(“关闭”);
	State string `json:"state"`
}
