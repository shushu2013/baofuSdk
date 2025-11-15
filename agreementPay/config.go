package agreementPay

import (
	"crypto/rsa"
	"fmt"

	"github.com/shushu2013/baofuSdk/tool"
)

type AgreementPayConfigParams struct {
	IsProdMode bool // true 生产，false 测试
	// 商户号（宝付提供）
	MemberId string `json:"member_id"`
	// 终端号（宝付提供）
	TerminalId string `json:"terminal_id"`

	// 商户私钥路径
	PrivateKeyPath string `json:"private_key_path"`
	// 商户私钥密码
	PrivateKeyPassword string `json:"private_key_password"`
	// 宝付公钥路径
	PublicKeyPath string `json:"public_key_path"`
}

func (c *AgreementPayConfigParams) validate() error {
	if c.MemberId == "" || c.TerminalId == "" ||
		c.PrivateKeyPath == "" || c.PrivateKeyPassword == "" ||
		c.PublicKeyPath == "" {
		return fmt.Errorf("配置缺少必填字段")
	}
	return nil
}

func (c *AgreementPayConfigParams) parseCert(config *AgreementPayConfig) error {
	// 从私钥文件中加载私钥
	privateKey, err := tool.GetPrivateKeyFromFile(c.PrivateKeyPath, c.PrivateKeyPassword)
	if err != nil {
		return err
	}
	config.PrivateKey = privateKey

	// 从公钥文件中加载公钥
	publicKey, err := tool.GetPublicKeyFromFile(c.PublicKeyPath)
	if err != nil {
		return err
	}
	config.PublicKey = publicKey

	return nil
}

// 宝财通3 接口配置
// https://docs.baofu.com/docs/bct3/baocaitong3
type AgreementPayConfig struct {
	IsProdMode bool // true 生产，false 测试
	// 商户号（宝付提供）
	MemberId string `json:"member_id"`
	// 终端号（宝付提供）
	TerminalId string `json:"terminal_id"`

	// 商户私钥 key
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	// 宝付公钥 key
	PublicKey *rsa.PublicKey `json:"public_key"`
}

func (c *AgreementPayConfig) GetBaseURL() string {
	baseUrl := BASE_AGREEMENT_PAY_API_URL
	if !c.IsProdMode {
		baseUrl = BASE_AGREEMENT_PAY_API_TEST_URL
	}

	return baseUrl
}

func NewAgreementPayConfig(config *AgreementPayConfigParams) (*AgreementPayConfig, error) {
	// 校验配置
	if err := config.validate(); err != nil {
		return nil, err
	}

	agreementPayConfig := &AgreementPayConfig{
		IsProdMode: config.IsProdMode,
		MemberId:   config.MemberId,
		TerminalId: config.TerminalId,
	}

	// 解析证书序列号
	if err := config.parseCert(agreementPayConfig); err != nil {
		return nil, err
	}

	return agreementPayConfig, nil
}
