package agreementPay

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/shushu2013/baofuSdk/tool"
)

func TestAgreementDirectPay(t *testing.T) {
	memberId := "102004459"   //商户号
	terminalId := "100005196" //终端号

	wd, _ := os.Getwd()
	pfxPath := path.Join(wd, "../cert", "BAOFU20240612_pri.pfx")     //商户私钥
	pubCerPath := path.Join(wd, "../cert", "BAOFUP20240612_pub.cer") //宝付公钥
	priKeyPass := "123456"

	configParams := &AgreementPayConfigParams{
		IsProdMode:         false,
		MemberId:           memberId,
		TerminalId:         terminalId,
		PrivateKeyPath:     pfxPath,
		PrivateKeyPassword: priKeyPass,
		PublicKeyPath:      pubCerPath,
	}
	config, err := NewAgreementPayConfig(configParams)
	if err != nil {
		t.Errorf("NewAgreementPayConfig failed: %v", err)
	}

	protocolNo := "1202411111040078480001616292"                          //签约协议号（确认绑卡返回）
	returnUrl := "http://10.0.60.66:8080/NotifyWeb/AgreeReturnUrl.action" //异步通知接收地址
	transId := "TID048127398CAA3310EA57"                                  //tool.CreateAeskey(20)

	cardInfo := "" //信用卡：信用卡有效期|安全码,借记卡：传空

	// 行业参数  (以下为游戏行业风控参，请参看接口文档附录风控参数)
	riskItem := map[string]string{
		"goodsCategory":    "06",                 //商品类目 详见附录《商品类目》
		"userLoginId":      "bofootest",          //用户在商户系统中的登陆名（手机号、邮箱等标识）
		"userEmail":        "",                   //用户邮箱
		"userMobile":       "15821798636",        //用户手机号
		"registerUserName": "大圣",                 //用户在商户系统中注册使用的名字
		"identifyState":    "1",                  //用户在平台是否已实名，1：是 ；0：不是
		"userIdNo":         "341182197807131732", //用户身份证号
		"registerTime":     "20170223113233",     //格式为：YYYYMMDDHHMMSS
		"registerIp":       "10.0.0.0",           //用户在商户端注册时留存的IP
		"chName":           "10.0.0.0",           //持卡人姓名
		"chIdNo":           "",                   //持卡人身份证号
		"chCardNo":         "",                   //持卡人银行卡号
		"chMobile":         "",                   //持卡人手机
		"chPayIp":          "116.216.217.170",    //持卡人支付IP
		"deviceOrderNo":    "",                   //加载设备指纹中的订单号

		/**--------行业参数  (以下为游戏行业风控参，请参看接口文档附录风控参数)-------------**/
		"gameName":       "15821798636", //充值游戏名称
		"userAcctId":     "15821798636", //游戏账户ID
		"rechargeType":   "0",           //充值类型 (0:为本账户充值或支付、1:为他人账户充值或支付； 默认为 0)
		"gameProdType":   "02",          //01：点券类 、 02：金币类 、 03：装备道具类 、 04：其他
		"gameAcctId":     "",            //被充值游戏账户ID,若充值类型为1 则填写
		"gameLoginTime":  "20",          //游戏登录次数，累计最近一个月
		"gameOnlineTime": "100",         //游戏在线时长，累计最近一个月
	}

	riskItemStr, _ := tool.StringifyJSON(riskItem)

	reqMap := &AgreementDirectPayRequest{
		MsgId:      tool.GetMsgId(),
		TransId:    transId,
		UserId:     "",
		ProtocolNo: protocolNo,
		TxnAmt:     "1",
		CardInfo:   cardInfo,
		RiskItem:   riskItemStr,
		ReturnUrl:  returnUrl,
	}

	resp, err := AgreementDirectPay(config, reqMap)
	if err != nil {
		t.Errorf("AgreementDirectPay failed: %v", err)
	}

	fmt.Println(resp)
}
