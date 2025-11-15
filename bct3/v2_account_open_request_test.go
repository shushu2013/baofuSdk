package bct3

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAccountOpenRequest(t *testing.T) {
	memberId := "102005245"   //商户号
	terminalId := "200005972" //终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     //商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") //宝付公钥
	priKeyPass := "123456"

	configParams := &BCT3ConfigParams{
		MemberId:           memberId,
		TerminalId:         terminalId,
		IsProdMode:         false,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewBCT3Config(configParams)
	if err != nil {
		t.Errorf("NewBCT3Config failed: %v", err)
	}

	// 登录号商户端用户唯一标识，商户自定义（不重复）
	loginNo := fmt.Sprintf("LN%s", tool.RandomStr(15))

	// 用户信息
	//企业开户
	email := "greatwalltesta@baofu.com"     //企业开户邮箱
	selfEmployed := false                   //是否个体户 企业为false，不传默认为false
	corporateName := "杨雪宇"                  //法人姓名
	corporateCertType := "ID"               //法人证件类型：身份证-ID，港澳通行证-HONG_KONG_AND_MACAO_PASS、台湾同胞来往内地通行证-TAIWAN_TRAVEL_PERMIT、护照-PASSPORT
	corporateCertId := "542624198102098035" //法人证件号码
	corporateMobile := "18756161713"        //法人手机号
	industryId := "9999"                    //行业分类代码，参考《行业分类代码表》

	//银行信息
	bankName := "中国工商银行"         //银行名称
	depositBankProvince := "江苏省" //开户行省份
	depositBankCity := "南京市"     //开户行城市
	depositBankName := "支行营业部"   //开户支行名称
	registerCapital := "50000"   //注册资本

	//个人开户
	mobileNo := "17127004795" //手机号

	//公共
	cardNo := "83522347826188"            //卡号 个人为对私卡号，企业为对公卡号
	customerName := "建国强文科技股份有限公司"        //客户名称,当为企业/个体时为企业名称
	certificateNo := "90540600JW0J70MKU1" //证件号，当为个人时上送身份证号，企业时上送营业执照号

	req := &AccOpenReq{
		AccType: ACCOUNT_TYPE_ENTERPRISE,
		AccInfo: AccOpenInfo{
			TransSerialNo:       tool.GetTransSerialNo(),
			LoginNo:             loginNo,
			MobileNo:            mobileNo,
			CardNo:              cardNo,
			CustomerName:        customerName,
			CertificateNo:       certificateNo,
			CertificateType:     CERT_TYPE_BUSINESS_LICENSE,
			Email:               email,
			SelfEmployed:        selfEmployed,
			CorporateName:       corporateName,
			CorporateCertType:   corporateCertType,
			CorporateCertId:     corporateCertId,
			CorporateMobile:     corporateMobile,
			IndustryId:          industryId,
			BankName:            bankName,
			DepositBankProvince: depositBankProvince,
			DepositBankCity:     depositBankCity,
			DepositBankName:     depositBankName,
			RegisterCapital:     registerCapital,
		},
		NoticeUrl: "https://www.baidu.com",
	}

	response := &AccOpenResp{}

	response, err = AccountOpenRequest(config, req)

	respStr, _ := tool.StringifyJSON(response)
	log.Printf("AccountOpenRequest response: %s", respStr)

	if err != nil {
		t.Errorf("AccountOpenRequest failed: %v", err)
	}
}
