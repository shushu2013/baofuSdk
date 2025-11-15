package tool

import (
	"fmt"
	"testing"
)

// 使用示例和测试
func TestSignatureUtils(t *testing.T) {

	fmt.Println("=== Go语言签名工具类示例 ===")

	// 示例数据
	encryptStr := `{"charset":"UTF-8","dgtlEnvlp":"23b88946e7af335f0a517e63a9c7dd95ef9f2ec4371ef0837206bf4fde29af2e4eaef92625bd5ed487bda578462cb4cdeee1322012146744c4f03a8632f0b8cb8162bd860212faa3f101e74c90b93359e101dbfd545624164dfb9b3e09518f62d93f154a65310806a452829f0919c5b28e7e5e5be5c15bf690810659c6c05416","memberId":"102005245","ncrptnSN":"00b9f3e90c370a7a0f","signSN":"00823078efb7e1bbdc","terminalId":"200005972","timestamp":"2025-10-27 11:26:37","verifyType":"11","version":"1.0"}{"accInfo":{"bankName":"中国工商银行","cardNo":"c9b286b6a05d88fc36f97f41469ba93c","certificateNo":"6e54db8c5da71265b4e163272dedb384905e599df4a3aef1bc1647a5da3dc4ca","certificateType":"LICENSE","corporateCertId":"881105294a68a713d649a695fd2fe7a6db6543bad256d19f28026606436cfe01","corporateCertType":"ID","corporateName":"杨雪宇","customerName":"建国强文科技股份有限公司","depositBankCity":"南京市","depositBankName":"支行营业部","depositBankProvince":"江苏省","email":"greatwalltesta@baofu.com","industryId":"9999","loginNo":"LN939338151248693","needUploadFile":false,"registerCapital":"50000","selfEmployed":false,"transSerialNo":"TSN8805DDDA8F81EA81F062"},"accType":2,"businessType":"BCT3","noticeUrl":"","version":"1.0.0"}`

	// 检查证书文件是否存在
	pfxPath := "/Users/wwb/项目文档/宝财通/宝付密钥/测试密钥/BAOFU20240612_pri.pfx"     //商户私钥
	pubCerPath := "/Users/wwb/项目文档/宝财通/宝付密钥/测试密钥/BAOFUP20240612_pub.cer" //宝付公钥
	priKeyPass := "123456"                                               //私钥密码

	// 1. 签名示例
	fmt.Println("\n1. 签名过程:")
	signature, err := encryptByRSA(encryptStr, pfxPath, priKeyPass, SIGNATURE_SHA1_WITH_RSA_ALGORITHM)
	if err != nil {
		fmt.Printf("签名失败: %v\n", err)
		return
	}
	fmt.Printf("原始数据: %s\n", encryptStr)
	fmt.Printf("签名结果: %s\n", signature)

	// 2. 验签示例
	fmt.Println("\n2. 验签过程:")
	encryptStr1 := `{"dgtlEnvlp":"1C96DF37898A9E5706C5D9C6A38A9E1F8A2DC299DFA2A52D4CA94811C70222FA2FE7023C4394C4D260ED6ADDB7CDC47FD82F36E58F9B8F2C9B23CC6467A3564398F1D82821143DA17BFFC05BB14A2B6E49C1460FE1E425797F746712B5A78A93E204D838B76363891EDA249CD8C80A9161F40734CB66879EADA3991F0CEF77ECE6EE3148B94C650D7569072E47EFA6598B0B37E65A829EA15E8A3AE68654B98AB4F151DCB9885552D974E1817D1129CA3451140D600D52E5989013653909D1BE44E9D19ACAE8FD81C1BCCCF7472EC7A41CB73F115F9DA485C01E88F27D96B374223002315FE811562A13B598B6FFD42A6EDA9C933F32C4D97E61ABB09C8ED98C","memberId":"102005245","ncrptnSN":"00823078efb7e1bbdc","serviceTp":"BCT3-1003-001-01","signSN":"00b9f3e90c370a7a0f","sysRespCode":"S_0000","sysRespDesc":"请求正常","terminalId":"200005972","verifyType":"11"}{"result":{"platformNo":"102005245","customerType":"2","contractNo":"CM610000000000275588","contractName":"涛强文科技股份有限公司","certificateNo":"4fca6b72f286d3457c8c25c77c2b19a07093fcde870e4393cd57a3b48795d337","customerName":"涛强文科技股份有限公司","email":"greatwalltesta@baofu.com","certificateType":"LICENSE"},"retCode":1,"version":"1.0.0"}`
	signature1 := "040B37AF952C61A2B227D90354BDCFC885AC8FB05A2EB425E215F98E3556157F53ACEE3DE534E1D439B39585FCB8449BEB54DB95FF017D5B857C6F1E0DD3176B557BE2051E63D24436CCA1DDFC7ACBB596491765D925EA4AAC90ED6C63043876D928FE726AEA18CAF128DA3F3A32C27D82C3AF1C2A775A2C013EC3433DBB18EA"
	isValid, err := verifySignature(pubCerPath, encryptStr1, signature1, SIGNATURE_SHA256_WITH_RSA_ALGORITHM)
	if err != nil {
		fmt.Printf("验签失败: %v\n", err)
		return
	}
	fmt.Printf("验签结果: %v\n", isValid)

	if isValid {
		fmt.Println("✅ 签名验证成功！")
	} else {
		fmt.Println("❌ 签名验证失败！")
	}
}
