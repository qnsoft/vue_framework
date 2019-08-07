package Kuaifutong

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"qnsoft/qn_web_api/utils/DateHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/php2go"

	"github.com/ajg/form"
	"golang.org/x/crypto/pkcs12"
)

type KuaiPayCore struct {
}

/*
参数集
*/
func (k *KuaiPayCore) SignParameters(parameters []map[string]string, encoding string) map[string]string {
	rt := map[string]string{}
	fmt.Println("待签名数据：")
	fmt.Println(parameters)
	prestr := k.CreateLinkString(k.ParamsFilter(parameters))
	value := k.Sign(prestr, CertPath, Key_In)
	for i := 0; i < len(parameters); i++ {
		for k, v := range parameters[i] {
			rt[k] = v
		}
	}
	rt["signatureAlgorithm"] = "RSA"
	rt["signatureInfo"] = value
	fmt.Println("排序后数据：")
	fmt.Println(parameters)
	fmt.Println(value)
	return rt
}

/*
参数处理
*/
func (k *KuaiPayCore) ParamsFilter(parameters []map[string]string) []map[string]string {
	//var sortedDictionary map[string]string
	for i := 0; i < len(parameters); i++ {
		for k, v := range parameters[i] {
			if k == "signatureAlgorithm" || k == "signatureInfo" || v == "" {
				parameters = append(parameters[:i])
			}
		}

	}
	fmt.Println(parameters)
	return parameters
}

/*
处理地址栏拼接字符串及排序
*/
func (k *KuaiPayCore) CreateLinkString(sortedParameters []map[string]string) string {
	stringBuilder := ""
	mapall := make(map[string]string)
	for i := 0; i < len(sortedParameters); i++ {
		for k, v := range sortedParameters[i] {
			mapall[k] = v
		}
	}
	var keys []string
	for k := range mapall {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		stringBuilder = stringBuilder + k + "=" + mapall[k] + "&"
	}
	//for i := 0; i < len(mapall); i++ {
	//for k, v := range mapall {
	//	stringBuilder = stringBuilder + k + "=" + v + "&"
	//}
	//}
	//去掉最后一个&
	_rt := php2go.Rtrim(stringBuilder, "&")
	//_rt := strings.Replace(stringBuilder, "1.0.0-PRD&", "1.0.0-PRD", 0)
	fmt.Println("排序后数据：")
	fmt.Println(_rt)
	return _rt
}

/*
签名
*/
func (k *KuaiPayCore) Sign(prestr, certFileName, password string) string {
	fmt.Println(prestr)
	signatureInfo := []byte(prestr)
	signer, _ := RsaSign(signatureInfo, certFileName, password)
	return php2go.Base64Encode(string(signer))

}

//私钥签名
func RsaSign(data []byte, certFileName, password string) ([]byte, error) {
	h := sha1.New()
	h.Write(data)
	hashed := h.Sum(nil)

	var pfxData []byte
	pfxData, err1 := ioutil.ReadFile(certFileName)
	if err1 != nil {
		return nil, err1
	}
	//获取私钥
	priv, _, err1 := pkcs12.Decode(pfxData, password)
	if err1 != nil {
		return nil, err1
	}
	private := priv.(*rsa.PrivateKey)
	return rsa.SignPKCS1v15(rand.Reader, private, crypto.SHA1, hashed)
}

type KuaiPayHelper struct {
}

/*
获取基础参数
*/
func (k *KuaiPayHelper) GetBaseParam() []map[string]string {
	var _rt []map[string]string
	//接口版本号 接口的版本号,便于在升级版本后,能新旧版本同时运行测试环境：1.0.0-IEST 生产环境：1.0.0-PRD
	_rt = append(_rt, map[string]string{"version": Version})
	//参数字符集 商户发送请求时,参数值使用的编码字符集，目前只允许UTF-8
	_rt = append(_rt, map[string]string{"charset": "UTF-8"})
	//语言 用于支持不同语言.java国际化中标准的locale值，如zh_CN表示中文，en表示英文；
	_rt = append(_rt, map[string]string{"language": "zh_CN"})
	//调用端IP
	_rt = append(_rt, map[string]string{"callerIp": Server_IP})
	return _rt
}
func (k *KuaiPayHelper) HttpPostCall(url string, formData map[string]string) string {
	var _rt = ""
	ErrorHelper.LogInfo("这是支付提交的URL：", url)
	ErrorHelper.LogInfo("这是支付提交的参数对象：", formData)

	_formData, err := form.EncodeToString(formData)
	ErrorHelper.CheckErr(err)
	if len(url) > 0 && len(formData) > 0 {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		payload := strings.NewReader(_formData)
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded; text/html; charset=utf-8")
		req.Header.Add("Accept", "text/html")
		res, err := client.Do(req)
		//res, err := http.DefaultClient.Do(req)
		ErrorHelper.CheckErr(err)
		defer res.Body.Close()
		if res != nil {
			body, _ := ioutil.ReadAll(res.Body)
			_rt = string(body)
		}
	}
	return _rt
}
func (k *KuaiPayHelper) OnRemoteCertificateValidationCallback() bool {
	return true
}

/*---------------------3.3快捷代扣接口------------------------------*/
/*
*3.3快捷代扣接口
*此接口用于为商户提供购买支付业务收款的接口.每次调用该接口会发送短信验证码到用户绑定的手机号。
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_sms_collect(inputDto Model_35) Return_33 {
	var _rt Return_33
	_baseParam := k.GetBaseParam()
	num := 0
	num2 := 0
	num3 := 0
	if num3 < num2 {
		num2 = num3
	}
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_sms_collect"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})     //商户身份ID
	_baseParam = append(_baseParam, map[string]string{"secMerchantId": R1_merchantNo})  //二级商户ID
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00004"})         //产品编号
	_baseParam = append(_baseParam, map[string]string{"orderNo": inputDto.OrderNo})     //订单编号
	_baseParam = append(_baseParam, map[string]string{"tradeName": inputDto.TradeName}) //交易名称
	//_baseParam = append(_baseParam, map[string]string{"merchantBankAccountNo": ""})                                //商户银行账号 可空（T+0模式必填）
	_baseParam = append(_baseParam, map[string]string{"tradeTime": date.FormatDate(time.Now(), "yyyyMMddHHmmss")}) //商户方交易时 此次交易在商户发发起的时间,注意此时间取值一般为商户方系统时间而非快付通生成此时间
	_baseParam = append(_baseParam, map[string]string{"amount": strconv.Itoa(num)})                                //交易金额 此次交易的具体金额,单位:分,不支持小数点
	_baseParam = append(_baseParam, map[string]string{"currency": "CNY"})                                          //币种
	_baseParam = append(_baseParam, map[string]string{"custBankNo": inputDto.CustBankNo})                          //客户银行账户行别
	//_baseParam = append(_baseParam, map[string]string{"custBankAccountIssuerNo": ""})                              //客户开户行网点号 指支付系统里的行号，具体到某个支行（网点）号
	_baseParam = append(_baseParam, map[string]string{"custBankAccountNo": inputDto.CustBankAccountNo}) //客户银行账户号
	_baseParam = append(_baseParam, map[string]string{"custName": inputDto.CustName})
	_baseParam = append(_baseParam, map[string]string{"custCertificationType": "0"})
	_baseParam = append(_baseParam, map[string]string{"custID": inputDto.CustID})
	_baseParam = append(_baseParam, map[string]string{"rateAmount": strconv.Itoa(num2)})
	_baseParam = append(_baseParam, map[string]string{"notifyUrl": QuickPay_Pay_notifyUrl})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.5 单笔付款接口------------------------------*/
/*
*3.5 单笔付款接口
*单笔付款在功能和接口参数上与单笔收款基本一致,只是付款方变成了商户自己,收款方变成了商户指定的客户
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_pay(_model Model_35) Return_35 {
	var _rt Return_35
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_pay"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})   //商户身份ID
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00002"})       //产品编号
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})     //订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
	_baseParam = append(_baseParam, map[string]string{"tradeName": _model.TradeName}) //交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
	if _model.MerchantBankAccountNo != "" {
		_baseParam = append(_baseParam, map[string]string{"merchantBankAccountNo": _model.MerchantBankAccountNo}) //商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
	}
	if _model.MerchantBindPhoneNo != "" {
		_baseParam = append(_baseParam, map[string]string{"merchantBindPhoneNo": _model.MerchantBindPhoneNo}) //商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
	}
	_baseParam = append(_baseParam, map[string]string{"tradeTime": date.FormatDate(time.Now(), "yyyyMMddHHmmss")}) //商户方交易时间 此次交易在商户发发起的时间,注意此时间取值一般为商户方系统时间而非快付通生成此时间
	_baseParam = append(_baseParam, map[string]string{"amount": _model.Amount})                                    //交易金额 此次交易的具体金额,单位:分,不支持小数点
	_baseParam = append(_baseParam, map[string]string{"currency": "CNY"})                                          //币种
	_baseParam = append(_baseParam, map[string]string{"custBankNo": _model.CustBankNo})                            //客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
	if _model.CustBankAccountIssuerNo != "" {
		_baseParam = append(_baseParam, map[string]string{"custBankAccountIssuerNo": _model.CustBankAccountIssuerNo}) //客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
	}
	_baseParam = append(_baseParam, map[string]string{"custBankAccountNo": _model.CustBankAccountNo}) //客户银行账户号 本次交易中,往客户的哪张卡上付钱
	_baseParam = append(_baseParam, map[string]string{"custName": _model.CustName})                   //客户姓名 收钱的客户的真实姓名
	if _model.CustBankAcctType != "" {
		_baseParam = append(_baseParam, map[string]string{"custBankAcctType": _model.CustBankAcctType}) //客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
	}
	if _model.CustAccountCreditOrDebit != "" {
		_baseParam = append(_baseParam, map[string]string{"custAccountCreditOrDebit": _model.CustAccountCreditOrDebit}) //客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
	}
	if _model.CustCardValidDate != "" {
		_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate}) //客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
	}
	if _model.CustCardCvv2 != "" {
		_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2}) //客户信用卡的cvv2 可空 信用卡的背面的三位数
	}
	_baseParam = append(_baseParam, map[string]string{"custCertificationType": "0"}) //客户证件类型 固定值0
	_baseParam = append(_baseParam, map[string]string{"custID": _model.CustID})      //客户证件号码
	if _model.CustPhone != "" {
		_baseParam = append(_baseParam, map[string]string{"custPhone": _model.CustPhone}) //客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
	}
	if _model.Messages != "" {
		_baseParam = append(_baseParam, map[string]string{"messages": _model.Messages}) //发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	}
	if _model.CustEmail != "" {
		_baseParam = append(_baseParam, map[string]string{"custEmail": _model.CustEmail}) //客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
	}
	if _model.EmailMessages != "" {
		_baseParam = append(_baseParam, map[string]string{"emailMessages": _model.EmailMessages}) //发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	}
	if _model.Remark != "" {
		_baseParam = append(_baseParam, map[string]string{"remark": _model.Remark}) //备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
	}
	if _model.CustProtocolNo != "" {
		_baseParam = append(_baseParam, map[string]string{"custProtocolNo": _model.CustProtocolNo}) //客户协议编号 可空 扣款人在快付通备案的协议号。
	}
	if _model.ExtendParams != "" {
		_baseParam = append(_baseParam, map[string]string{"extendParams": _model.ExtendParams}) //扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
	}
	_baseParam = append(_baseParam, map[string]string{"rateAmount": _model.RateAmount})     //商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
	_baseParam = append(_baseParam, map[string]string{"notifyUrl": QuickPay_Pay_notifyUrl}) //商户后台通知URL 可空 当交易完成后，快付通会URL异步通知商户(只有非终态才有回调) 回调请求参数：orderNo=9897013867 amount=100 merchantId=2018041100 status=1 failureDetails=成功 errorCode=xxxx
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*
*3.5 单笔付款接口(小额)
*单笔付款在功能和接口参数上与单笔收款基本一致,只是付款方变成了商户自己,收款方变成了商户指定的客户
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_pay_small(_model Model_35) Return_35 {
	var _rt Return_35
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_pay"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R3_merchantNo})   //商户身份ID
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00002"})       //产品编号
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})     //订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
	_baseParam = append(_baseParam, map[string]string{"tradeName": _model.TradeName}) //交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
	if _model.MerchantBankAccountNo != "" {
		_baseParam = append(_baseParam, map[string]string{"merchantBankAccountNo": _model.MerchantBankAccountNo}) //商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
	}
	if _model.MerchantBindPhoneNo != "" {
		_baseParam = append(_baseParam, map[string]string{"merchantBindPhoneNo": _model.MerchantBindPhoneNo}) //商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
	}
	_baseParam = append(_baseParam, map[string]string{"tradeTime": date.FormatDate(time.Now(), "yyyyMMddHHmmss")}) //商户方交易时间 此次交易在商户发发起的时间,注意此时间取值一般为商户方系统时间而非快付通生成此时间
	_baseParam = append(_baseParam, map[string]string{"amount": _model.Amount})                                    //交易金额 此次交易的具体金额,单位:分,不支持小数点
	_baseParam = append(_baseParam, map[string]string{"currency": "CNY"})                                          //币种
	_baseParam = append(_baseParam, map[string]string{"custBankNo": _model.CustBankNo})                            //客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
	if _model.CustBankAccountIssuerNo != "" {
		_baseParam = append(_baseParam, map[string]string{"custBankAccountIssuerNo": _model.CustBankAccountIssuerNo}) //客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
	}
	_baseParam = append(_baseParam, map[string]string{"custBankAccountNo": _model.CustBankAccountNo}) //客户银行账户号 本次交易中,往客户的哪张卡上付钱
	_baseParam = append(_baseParam, map[string]string{"custName": _model.CustName})                   //客户姓名 收钱的客户的真实姓名
	if _model.CustBankAcctType != "" {
		_baseParam = append(_baseParam, map[string]string{"custBankAcctType": _model.CustBankAcctType}) //客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
	}
	if _model.CustAccountCreditOrDebit != "" {
		_baseParam = append(_baseParam, map[string]string{"custAccountCreditOrDebit": _model.CustAccountCreditOrDebit}) //客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
	}
	if _model.CustCardValidDate != "" {
		_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate}) //客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
	}
	if _model.CustCardCvv2 != "" {
		_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2}) //客户信用卡的cvv2 可空 信用卡的背面的三位数
	}
	_baseParam = append(_baseParam, map[string]string{"custCertificationType": "0"}) //客户证件类型 固定值0
	_baseParam = append(_baseParam, map[string]string{"custID": _model.CustID})      //客户证件号码
	if _model.CustPhone != "" {
		_baseParam = append(_baseParam, map[string]string{"custPhone": _model.CustPhone}) //客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
	}
	if _model.Messages != "" {
		_baseParam = append(_baseParam, map[string]string{"messages": _model.Messages}) //发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	}
	if _model.CustEmail != "" {
		_baseParam = append(_baseParam, map[string]string{"custEmail": _model.CustEmail}) //客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
	}
	if _model.EmailMessages != "" {
		_baseParam = append(_baseParam, map[string]string{"emailMessages": _model.EmailMessages}) //发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	}
	if _model.Remark != "" {
		_baseParam = append(_baseParam, map[string]string{"remark": _model.Remark}) //备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
	}
	if _model.CustProtocolNo != "" {
		_baseParam = append(_baseParam, map[string]string{"custProtocolNo": _model.CustProtocolNo}) //客户协议编号 可空 扣款人在快付通备案的协议号。
	}
	if _model.ExtendParams != "" {
		_baseParam = append(_baseParam, map[string]string{"extendParams": _model.ExtendParams}) //扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
	}
	_baseParam = append(_baseParam, map[string]string{"rateAmount": _model.RateAmount})     //商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
	_baseParam = append(_baseParam, map[string]string{"notifyUrl": QuickPay_Pay_notifyUrl}) //商户后台通知URL 可空 当交易完成后，快付通会URL异步通知商户(只有非终态才有回调) 回调请求参数：orderNo=9897013867 amount=100 merchantId=2018041100 status=1 failureDetails=成功 errorCode=xxxx
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.6快捷协议代扣协议申请接口------------------------------*/
/*
*3.6快捷协议代扣协议申请接口
*此接口用于商户平台申请开通快捷协议代扣，并由快付通发送手机验证码。
 */
func (k *KuaiPayHelper) Gbp_same_id_treaty_collect_apply(_model Model_36) Return_36 {
	var _rt Return_36
	_baseParam := k.GetBaseParam()
	//订单编号
	orderNumber := date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99))
	fmt.Println(orderNumber)
	_baseParam = append(_baseParam, map[string]string{"orderNo": orderNumber})
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_treaty_collect_apply"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"startDate": _model.StartDate})
	if _model.Note != "" { //说明 参数可空
		_baseParam = append(_baseParam, map[string]string{"note": _model.Note})
	}
	_baseParam = append(_baseParam, map[string]string{"holderName": _model.HolderName})
	_baseParam = append(_baseParam, map[string]string{"bankType": _model.BankType})
	_baseParam = append(_baseParam, map[string]string{"bankCardType": _model.BankCardType})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": _model.BankCardNo})
	_baseParam = append(_baseParam, map[string]string{"mobileNo": _model.MobileNo})
	_baseParam = append(_baseParam, map[string]string{"certificateType": _model.CertificateType})
	_baseParam = append(_baseParam, map[string]string{"certificateNo": _model.CertificateNo})
	flag := false
	if _model.BankCardType == "1" { //借记卡
		flag = true
	}
	if flag {
		_baseParam = append(_baseParam, map[string]string{"treatyType": _model.TreatyType}) //协议类型 11：借记卡扣款 12：信用卡扣款
		_baseParam = append(_baseParam, map[string]string{"endDate": _model.EndDate})       //协议失效日期 如果是借记卡 协议失效日是填写日期
	} else {
		if _model.BankCardType == "2" { //信用卡
			_baseParam = append(_baseParam, map[string]string{"treatyType": _model.TreatyType})               //协议类型 11：借记卡扣款 12：信用卡扣款
			_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate}) //客户信用卡有效期
			_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2})           //客户信用卡的cvv2
			_baseParam = append(_baseParam, map[string]string{"endDate": _model.EndDate})                     //如果信用卡 协议失效日即信用卡失效日
		}
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*
*3.6快捷协议代扣协议申请接口(小额专用)
*此接口用于商户平台申请开通快捷协议代扣，并由快付通发送手机验证码。
 */
func (k *KuaiPayHelper) Gbp_same_id_treaty_collect_apply_small(_model Model_36) Return_36 {
	var _rt Return_36
	_baseParam := k.GetBaseParam()
	//订单编号
	orderNumber := date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99))
	fmt.Println(orderNumber)
	_baseParam = append(_baseParam, map[string]string{"orderNo": orderNumber})
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_treaty_collect_apply"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R3_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"startDate": _model.StartDate})
	if _model.Note != "" { //说明 参数可空
		_baseParam = append(_baseParam, map[string]string{"note": _model.Note})
	}
	_baseParam = append(_baseParam, map[string]string{"holderName": _model.HolderName})
	_baseParam = append(_baseParam, map[string]string{"bankType": _model.BankType})
	_baseParam = append(_baseParam, map[string]string{"bankCardType": _model.BankCardType})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": _model.BankCardNo})
	_baseParam = append(_baseParam, map[string]string{"mobileNo": _model.MobileNo})
	_baseParam = append(_baseParam, map[string]string{"certificateType": _model.CertificateType})
	_baseParam = append(_baseParam, map[string]string{"certificateNo": _model.CertificateNo})
	flag := false
	if _model.BankCardType == "1" { //借记卡
		flag = true
	}
	if flag {
		_baseParam = append(_baseParam, map[string]string{"treatyType": _model.TreatyType}) //协议类型 11：借记卡扣款 12：信用卡扣款
		_baseParam = append(_baseParam, map[string]string{"endDate": _model.EndDate})       //协议失效日期 如果是借记卡 协议失效日是填写日期
	} else {
		if _model.BankCardType == "2" { //信用卡
			_baseParam = append(_baseParam, map[string]string{"treatyType": _model.TreatyType})               //协议类型 11：借记卡扣款 12：信用卡扣款
			_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate}) //客户信用卡有效期
			_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2})           //客户信用卡的cvv2
			_baseParam = append(_baseParam, map[string]string{"endDate": _model.EndDate})                     //如果信用卡 协议失效日即信用卡失效日
		}
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.7快捷协议代扣协议确定接口------------------------------*/
/*
*3.7快捷协议代扣协议确定接口
*商户平台通过此接口确认开通代扣协议，进行四要素鉴权获取进行快捷协议代扣的协议号
@model_37 对象
@TreatyType 协议类型
*/
func (k *KuaiPayHelper) Gbp_same_id_confirm_treaty_collect_apply(_model Model_37, TreatyType string) Return_37 {
	var _rt Return_37
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_confirm_treaty_collect_apply"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"smsSeq": _model.SmsSeq})
	_baseParam = append(_baseParam, map[string]string{"authCode": _model.AuthCode})
	_baseParam = append(_baseParam, map[string]string{"holderName": _model.HolderName})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": _model.BankCardNo})
	//	flag := false
	if TreatyType == "12" {
		_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate})
		_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2})
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*
*3.7快捷协议代扣协议确定接口(小额专用)
*商户平台通过此接口确认开通代扣协议，进行四要素鉴权获取进行快捷协议代扣的协议号
@model_37 对象
@TreatyType 协议类型
*/
func (k *KuaiPayHelper) Gbp_same_id_confirm_treaty_collect_apply_small(_model Model_37, TreatyType string) Return_37 {
	var _rt Return_37
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_confirm_treaty_collect_apply"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R3_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"smsSeq": _model.SmsSeq})
	_baseParam = append(_baseParam, map[string]string{"authCode": _model.AuthCode})
	_baseParam = append(_baseParam, map[string]string{"holderName": _model.HolderName})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": _model.BankCardNo})
	//	flag := false
	if TreatyType == "12" {
		_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate})
		_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2})
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.8快捷协议代扣接口------------------------------*/
/*
*3.8快捷协议代扣接口
*此接口用于商户平台协议代扣。此接口需要用户先签定协议。
@_model 38实体
@TreatyType协议类型
*/
func (k *KuaiPayHelper) Gbp_same_id_credit_card_treaty_collect(_model Model_38, TreatyType, _merchantNo string) Return_38 {
	var _rt Return_38
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_treaty_collect"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	_baseParam = append(_baseParam, map[string]string{"merchantId": _merchantNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00004"})
	_baseParam = append(_baseParam, map[string]string{"treatyNo": _model.TreatyNo})
	_baseParam = append(_baseParam, map[string]string{"tradeTime": _model.TradeTime})
	_baseParam = append(_baseParam, map[string]string{"amount": _model.Amount})
	_baseParam = append(_baseParam, map[string]string{"currency": "CNY"})
	_baseParam = append(_baseParam, map[string]string{"rateAmount": _model.RateAmount})
	_baseParam = append(_baseParam, map[string]string{"holderName": _model.HolderName})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": _model.BankCardNo})
	_baseParam = append(_baseParam, map[string]string{"bankType": _model.BankType})
	if TreatyType == "12" || TreatyType == "15" {
		_baseParam = append(_baseParam, map[string]string{"custCardValidDate": _model.CustCardValidDate})
		_baseParam = append(_baseParam, map[string]string{"custCardCvv2": _model.CustCardCvv2})
	}
	_baseParam = append(_baseParam, map[string]string{"notifyUrl": Delegate_Pay_notifyUrl})
	if _model.CityCode != "" {
		_baseParam = append(_baseParam, map[string]string{"cityCode": _model.CityCode})
	}
	if _model.SourceIP != "" {
		_baseParam = append(_baseParam, map[string]string{"sourceIP": _model.SourceIP})
	}
	if _model.DeviceID != "" {
		_baseParam = append(_baseParam, map[string]string{"deviceID": _model.DeviceID})
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.9快捷协议代扣协议查询接口------------------------------*/
/*
*3.9快捷协议代扣协议查询接口
*此接口用于商户平台通过此接口查询协议信息。
 */
func (k *KuaiPayHelper) Gbp_query_treaty_info(_model Model_39) Return_39 {
	var _rt Return_39
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_query_treaty_info"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	if _model.TreatyNo == "" {
		_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	}
	if _model.OrderNo == "" {
		_baseParam = append(_baseParam, map[string]string{"treatyNo": _model.TreatyNo})
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.10快捷协议代扣协议解除接口------------------------------*/
/*
*3.10快捷协议代扣协议解除接口
*此接口用于商户平台通过此接口解除快捷协议收款协议信息。
 */
func (k *KuaiPayHelper) Gbp_cancel_treaty_info(_model Model_310) Return_310 {
	var _rt Return_310
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_cancel_treaty_info"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"treatyNo": _model.TreatyNo})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.11用户资金查询接口------------------------------*/
/*
*3.11用户资金查询接口
*用于查询用户未付或者未扣资金。具体查询未付资金或者未扣资金，根据请求的账户编号所开通的业务类型而定。
*productNo 查询交易也是一种产品；具体使用的查询类产品编号,由快付通提供给商户；
*pageNum当前查询第几页的数据（不填默认1，每页200记录）
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_not_pay_balance(_model Model_311) Return_311 {
	var _rt Return_311
	if _model.PageNum == "" {
		_model.PageNum = "1"
	}
	_baseParam := k.GetBaseParam()
	if _model.ReqNo != "" {
		_baseParam = append(_baseParam, map[string]string{"reqNo": _model.ReqNo}) //请求编号(可空)
	}
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_not_pay_balance"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo}) //商户身份ID //
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00006"})     //产品编号
	if _model.CustID != "" {
		_baseParam = append(_baseParam, map[string]string{"custID": _model.CustID}) //身份证号
	}
	_baseParam = append(_baseParam, map[string]string{"pageNum": _model.PageNum}) //页码
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*
*3.11用户资金查询接口(小额专用)
*用于查询用户未付或者未扣资金。具体查询未付资金或者未扣资金，根据请求的账户编号所开通的业务类型而定。
*productNo 查询交易也是一种产品；具体使用的查询类产品编号,由快付通提供给商户；
*pageNum当前查询第几页的数据（不填默认1，每页200记录）
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_not_pay_balance_small(_model Model_311) Return_311 {
	var _rt Return_311
	if _model.PageNum == "" {
		_model.PageNum = "1"
	}
	_baseParam := k.GetBaseParam()
	if _model.ReqNo != "" {
		_baseParam = append(_baseParam, map[string]string{"reqNo": _model.ReqNo}) //请求编号(可空)
	}
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_not_pay_balance"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R3_merchantNo}) //商户身份ID //
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00006"})     //产品编号
	if _model.CustID != "" {
		_baseParam = append(_baseParam, map[string]string{"custID": _model.CustID}) //身份证号
	}
	_baseParam = append(_baseParam, map[string]string{"pageNum": _model.PageNum}) //页码
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//接下来将单据写入数据库
	return _rt
}

/*---------------------3.12交易查询接口------------------------------*/
/*
* 3.12交易查询接口
*用于查询指定的一笔或多笔交易的结果,例如购买支付交易状态
 */
func (k *KuaiPayHelper) Gbp_same_id_credit_card_trade_record_query(_model Model_312) Return_312 {
	var _rt Return_312
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_trade_record_query"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00005"})
	_baseParam = append(_baseParam, map[string]string{"startDate": _model.StartDate})
	_baseParam = append(_baseParam, map[string]string{"endDate": _model.EndDate})
	if _model.OrderNo != "" {
		_baseParam = append(_baseParam, map[string]string{"orderNo": _model.OrderNo})
	}
	if _model.TradeType != "" {
		_baseParam = append(_baseParam, map[string]string{"tradeType": _model.TradeType})
	}
	if _model.Status != "" {
		_baseParam = append(_baseParam, map[string]string{"status": _model.Status})
	}
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*-----------------------3.13账户余额查询接口-----------------------------*/
/*
*3.13账户余额查询接口
此功能用于给商户查询该商户的账户余额。
*/
func (k *KuaiPayHelper) Query_available_balance() Return_313 {
	var _rt Return_313
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "query_available_balance"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R2_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "2GCB0AAA"})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//准备写入数据库
	return _rt
}

/*-----------------------3.14交易类对账文件获取接口-----------------------------*/
/*
*3.14交易类对账文件获取接口
此功能用于给商户提供对账数据。
*/
func (k *KuaiPayHelper) Gbp_same_id_credit_card_recon_result_query(_model Model_314) Return_314 {
	var _rt Return_314
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_recon_result_query"})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00007"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"checkDate": _model.CheckDate})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//准备写入数据库
	return _rt
}

/*---------------------3.15银行卡三要素验证接口------------------------------*/
/*
*3.15银行卡三要素验证接口
*此接口用于校验指定的银行卡和用户身份信息是否匹配及正确，
*含二要素验证（户名、卡号）与三要素验证（户名、卡号、证件号）两种鉴权模式
 */
func (k *KuaiPayHelper) Gbp_threeMessage_verification(_model Model_315, orderStarWith string) Return_315 {
	var _rt Return_315
	//订单编号
	if orderStarWith == "" {
		orderStarWith = "KFT"
	}
	_orderNo := orderStarWith + date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99))
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_threeMessage_verification"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _orderNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00001"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"custBankNo": _model.CustBankNo})
	_baseParam = append(_baseParam, map[string]string{"custName": _model.CustName})
	if _model.CustBankAccountNo != "" {
		_baseParam = append(_baseParam, map[string]string{"custBankAccountNo": _model.CustBankAccountNo})
	}
	if _model.CustAccountCreditOrDebit != "" {
		_baseParam = append(_baseParam, map[string]string{"custAccountCreditOrDebit": _model.CustAccountCreditOrDebit})
	}
	_baseParam = append(_baseParam, map[string]string{"custCertificationType": "0"})
	_baseParam = append(_baseParam, map[string]string{"custID": _model.CustID})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//准备写入数据库
	return _rt
}

/*
3.16验证类对账文件获取接口
此功能用于给商户提供对账数据。
*/
func (k *KuaiPayHelper) Gbp_same_id_credit_card_verify_result_query(_model Model_316) Return_316 {
	var _rt Return_316
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_same_id_credit_card_verify_result_query"})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00007"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R1_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"checkDate": _model.CheckDate})

	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//准备写入数据库
	return _rt
}

/*---------------------3.17单笔付款接口(商户提现)------------------------------*/
/*
*3.17单笔付款接口
*单笔付款在功能和接口参数上与单笔收款基本一致,只是付款方变成了商户自己,收款方变成了商户指定的客户*
 */
func (k *KuaiPayHelper) Gbp_pay(Money float64) Return_317 {
	var _rt Return_317
	//订单编号
	orderNumber := date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99))
	fmt.Println(orderNumber)
	num := int(Money * 100)
	_baseParam := k.GetBaseParam()
	fmt.Println(_baseParam)
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_pay"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R2_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "GBP00010"})
	_baseParam = append(_baseParam, map[string]string{"orderNo": orderNumber})
	_baseParam = append(_baseParam, map[string]string{"tradeName": "商户提现"})
	_baseParam = append(_baseParam, map[string]string{"tradeTime": date.FormatDate(time.Now(), "yyMMddHHmmss")})
	_baseParam = append(_baseParam, map[string]string{"currency": "CNY"})
	_baseParam = append(_baseParam, map[string]string{"custProtocolNo": "30000043360199"})
	_baseParam = append(_baseParam, map[string]string{"custBankAccountNo": "6226223001466633"})
	_baseParam = append(_baseParam, map[string]string{"custBankNo": "3051000"})
	_baseParam = append(_baseParam, map[string]string{"custName": "李金领"})
	_baseParam = append(_baseParam, map[string]string{"amount": strconv.Itoa(num)})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*---------------------3.18单笔付款协议申请接口------------------------------*/
/*
*3.18单笔付款协议申请接口
*此接口用于平台对商户通过网络申请单笔收款的协议，方便商户提前备案协议。
*此接口返回的协议号，用于单笔收款接口代扣时验证客户信息。
 */
func (k *KuaiPayHelper) Gbp_send_treaty_record_to_kft() string {
	var _rt = ""
	_orderNo := date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99))
	_text := "KFT" + _orderNo
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_send_treaty_record_to_kft"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R2_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"orderNo": _orderNo})
	_baseParam = append(_baseParam, map[string]string{"merchantTreatyNo": _text})
	_baseParam = append(_baseParam, map[string]string{"treatyType": "2"})
	_baseParam = append(_baseParam, map[string]string{"paymentItem": "04903"})
	_baseParam = append(_baseParam, map[string]string{"startDate": date.FormatDate(time.Now(), "yyMMdd")})
	_baseParam = append(_baseParam, map[string]string{"endDate": date.FormatDate(time.Now(), "20201231")})
	_baseParam = append(_baseParam, map[string]string{"bankCardType": "1"})
	_baseParam = append(_baseParam, map[string]string{"holderName": "李金领"})
	_baseParam = append(_baseParam, map[string]string{"bankType": "3051000"})
	_baseParam = append(_baseParam, map[string]string{"bankCardNo": "6226223001466633"})
	_baseParam = append(_baseParam, map[string]string{"mobileNo": "15638905677"})
	_baseParam = append(_baseParam, map[string]string{"certificateType": "0"})
	_baseParam = append(_baseParam, map[string]string{"certificateNo": "41012219710510305X"})
	_baseParam = append(_baseParam, map[string]string{"currencyType": "CNY"})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}

/*---------------------3.19结算户对接接口------------------------------*/
/*
*3.19单笔付款协议查询接口
*此接口用于平台对商户通过网络线上查询单笔收款协议状态。
 */
func (k *KuaiPayHelper) Gbp_query_treaty_record_info(orderNumber string) string {
	var _rt = ""
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "gbp_query_treaty_record_info"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R2_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"orderNo": orderNumber})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	//准备写入数据库
	return _rt
}

/*
*3.22结算查询账户余额接口
 */
func (k *KuaiPayHelper) Query_available_balance_A() Return_322 {
	var _rt Return_322
	_baseParam := k.GetBaseParam()
	_baseParam = append(_baseParam, map[string]string{"service": "query_available_balance"})
	_baseParam = append(_baseParam, map[string]string{"merchantId": R2_merchantNo})
	_baseParam = append(_baseParam, map[string]string{"productNo": "2GCB0AAA"})
	_core := KuaiPayCore{}
	ErrorHelper.LogInfo(_core)
	formData := _core.SignParameters(_baseParam, "utf-8")
	fmt.Println(formData)
	_req := k.HttpPostCall(BaseUrl, formData)
	ErrorHelper.LogInfo("结算余额查询：", _req)
	err := json.Unmarshal([]byte(_req), &_rt)
	ErrorHelper.CheckErr(err)
	return _rt
}
