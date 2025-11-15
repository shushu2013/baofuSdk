package bct3

import (
	"crypto/rsa"
	"fmt"

	"github.com/shushu2013/baofuSdk/tool"
)

type BCT3ConfigParams struct {
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

func (c *BCT3ConfigParams) validate() error {
	if c.MemberId == "" || c.TerminalId == "" ||
		c.PrivateKeyPath == "" || c.PrivateKeyPassword == "" ||
		c.PublicKeyPath == "" {
		return fmt.Errorf("配置缺少必填字段")
	}
	return nil
}

func (c *BCT3ConfigParams) parseCert(config *BCT3Config) error {
	// 从私钥文件中提取序列号
	signSN, err := tool.GetPrivateKeySNFromFile(c.PrivateKeyPath, c.PrivateKeyPassword)
	if err != nil {
		return err
	}
	config.SignSN = signSN

	// 从公钥文件中提取序列号
	ncrptnSN, err := tool.GetPublicKeySNFromFile(c.PublicKeyPath)
	if err != nil {
		return err
	}
	config.NcrptnSN = ncrptnSN

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
type BCT3Config struct {
	IsProdMode bool // true 生产，false 测试
	// 商户号（宝付提供）
	MemberId string `json:"member_id"`
	// 终端号（宝付提供）
	TerminalId string `json:"terminal_id"`

	// 加密方式
	VerifyType string `json:"verify_type"`
	// 字符集(固定字符集UTF-8)
	Charset string `json:"charset"`
	// 接口版本号(固定1.0)
	Version string `json:"version"`

	// 商户私钥 key
	PrivateKey *rsa.PrivateKey `json:"private_key"`
	// 宝付公钥 key
	PublicKey *rsa.PublicKey `json:"public_key"`

	// 证书序列号（从证书中提取）
	SignSN   string `json:"sign_sn"`   // 签名证书序列号（商户私钥序列号）
	NcrptnSN string `json:"ncrptn_sn"` // 加密证书序列号（宝付公钥序列号）
}

func (c *BCT3Config) GetBaseURL(serviceTp string) string {
	baseUrl := BASE_BCT3_API_URL
	if !c.IsProdMode {
		baseUrl = BASE_BCT3_API_TEST_URL
	}

	return fmt.Sprintf("%s/%s/transReq.do", baseUrl, serviceTp)
}

func NewBCT3Config(config *BCT3ConfigParams) (*BCT3Config, error) {
	// 校验配置
	if err := config.validate(); err != nil {
		return nil, err
	}

	bct3Config := &BCT3Config{
		IsProdMode: config.IsProdMode,
		MemberId:   config.MemberId,
		TerminalId: config.TerminalId,
		VerifyType: VERIFY_TYPE_RSA,
		Charset:    "UTF-8",
		Version:    "1.0",
	}

	// 解析证书序列号
	if err := config.parseCert(bct3Config); err != nil {
		return nil, err
	}

	return bct3Config, nil
}
