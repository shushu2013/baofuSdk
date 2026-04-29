package refund

import (
	"crypto/rsa"
	"fmt"

	"github.com/shushu2013/baofuSdk/tool"
)

// RefundConfigParams 退款配置参数
type RefundConfigParams struct {
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

func (c *RefundConfigParams) validate() error {
	if c.MemberId == "" || c.TerminalId == "" ||
		c.PrivateKeyPath == "" || c.PrivateKeyPassword == "" ||
		c.PublicKeyPath == "" {
		return fmt.Errorf("配置缺少必填字段")
	}
	return nil
}

func (c *RefundConfigParams) parseCert(config *RefundConfig) error {
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

// RefundConfig 退款接口配置
type RefundConfig struct {
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

func (c *RefundConfig) GetBaseURL() string {
	if c.IsProdMode {
		return BASE_REFUND_API_URL
	}
	return BASE_REFUND_API_TEST_URL
}

// NewRefundConfig 创建退款配置
func NewRefundConfig(config *RefundConfigParams) (*RefundConfig, error) {
	// 校验配置
	if err := config.validate(); err != nil {
		return nil, err
	}

	refundConfig := &RefundConfig{
		IsProdMode: config.IsProdMode,
		MemberId:   config.MemberId,
		TerminalId: config.TerminalId,
	}

	// 解析证书
	if err := config.parseCert(refundConfig); err != nil {
		return nil, err
	}

	return refundConfig, nil
}

// RefundClient 退款客户端
type RefundClient struct {
	config *RefundConfig
}

// NewRefundClient 创建退款客户端
func NewRefundClient(config *RefundConfig) *RefundClient {
	return &RefundClient{config: config}
}
