package Kuaifutong

const (
	//(账户编号)河南联九舟网络科技有限公司
	R1_merchantNo string = "2018112302640260"
	//(账户编号)河南联九舟网络科技有限公司（结算账户）
	R2_merchantNo string = "2018112602642698"
	//(账户编号)小额业务专用
	R3_merchantNo string = "2019010902719239"
	//基础接口地址 https://merchant.kftpay.com.cn:8443/gateway/nonbatch 生产请求地址
	BaseUrl string = "https://merchant.kftpay.com.cn:8443/gateway/nonbatch"
	Key_In  string = "hnlj1125A"
	//证件扫描件路径
	CertPath string = "cer/zfb789.pfx"
	//接口版本号 测试环境：1.0.0-IEST 生产环境：1.0.0-PRD
	Version string = "1.0.0-PRD"
	//Version string = "1.0.0-IEST"
	//web基础地址
	WebUrl    string = "https://api1.xhdncppf.com"
	Server_IP        = "39.97.111.217"
	//Server_IP                     = "218.17.35.123"
	QuickPay_Pay_notifyUrl string = WebUrl + "/OnlinePay/kft/notify_kft"
	Delegate_Pay_notifyUrl string = WebUrl + "/OnlinePay/kft/notify_kft"
)

type PolicyType int32

const (
	DebCard    PolicyType = 11
	CreditCard PolicyType = 12
)

/*
*单笔付款接口
 */
type Return_5 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//列表详情
	Details string `json:"details"`
	//总共页数
	TotalPage string `json:"totalPage"`
	//当前页码
	PageNum string `json:"pageNum"`
	//每页记录数
	PageSize string `json:"pageSize"`
	//总记录数
	TotalCount string `json:"totalCount"`
}

/*
*3.3快捷代扣接口
 */
type Return_33 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//短信流水号
	SmsSeq string `json:"smsSeq"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.5单笔付款接口传参
 */
type Model_35 struct {
	//订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
	OrderNo string
	//交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
	TradeName string
	//商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
	MerchantBankAccountNo string
	//商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
	MerchantBindPhoneNo string
	//交易金额 此次交易的具体金额,单位:分,不支持小数点
	Amount string
	//客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
	CustBankNo string
	//客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
	CustBankAccountIssuerNo string
	//客户银行账户号 本次交易中,往客户的哪张卡上付钱
	CustBankAccountNo string
	//客户姓名 收钱的客户的真实姓名
	CustName string
	//客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
	CustBankAcctType string
	//客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
	CustAccountCreditOrDebit string
	//客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
	CustCardValidDate string
	//客户信用卡的cvv2 可空 信用卡的背面的三位数
	CustCardCvv2 string
	//客户证件号码
	CustID string
	//客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
	CustPhone string
	//发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	Messages string
	//客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
	CustEmail string
	//发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
	EmailMessages string
	//备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
	Remark string
	//客户协议编号 可空 扣款人在快付通备案的协议号。
	CustProtocolNo string
	//扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
	ExtendParams string
	//商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
	RateAmount string
}

/*
*3.5单笔付款接口返回
 */
type Return_35 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//银行返回时间
	BankReturnTime string `json:"bankReturnTime"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.6快捷协议代扣协议申请接口
 */
type Model_36 struct {
	//协议类型  商户申请开通的协议类型 //11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款
	TreatyType string
	//说明  商户开通协议的相关说明
	Note string
	//协议生效日期
	StartDate string
	//协议失效日期
	EndDate string
	//持卡人真实姓名
	HolderName string
	//银行行别
	BankType string
	//银行卡类型
	BankCardType string
	//银行卡号
	BankCardNo string
	//预留手机号码
	MobileNo string
	//证件类型
	CertificateType string
	//证件号
	CertificateNo string
	//客户信用卡有效期
	CustCardValidDate string
	//客户信用卡的cvv2
	CustCardCvv2 string
}

/*
*3.6快捷协议代扣协议申请接口返回对象
 */
type Return_36 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//短信流水号
	SmsSeq string `json:"smsSeq"`
	//交易状态
	Status string `json:"status"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.7快捷协议代扣协议确定接口
 */
type Model_37 struct {
	//订单编号
	OrderNo string
	//短信流水号
	SmsSeq string
	//手机动态校验码
	AuthCode string
	//持卡人真实姓名
	HolderName string
	//银行卡号
	BankCardNo string
	//客户信用卡有效期
	CustCardValidDate string
	//客户信用卡的cvv2
	CustCardCvv2 string
}

/*
*3.7快捷协议代扣协议确定接口返回对象
 */
type Return_37 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//协议号 可空
	TreatyId string `json:"treatyId"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.8快捷协议代扣接口
 */
type Model_38 struct {
	//订单编号
	OrderNo string
	//协议号
	TreatyNo string
	//商户方交易时间
	TradeTime string
	//代扣金额
	Amount string
	//账户ID 可空 快付通的客户账户号 协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：余额+借记卡扣款 15： 余额+信用卡扣款
	CustAccountId string
	//持卡人真实姓名 可空 协议类型：协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：	余额+借记卡扣款 15： 余额+信用卡扣款 协议代扣申请接口时，如果协议类型为11、12、14、15时不可为空
	HolderName string
	//银行行别 可空 协议类型：协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：	余额+借记卡扣款 15： 余额+信用卡扣款协议代扣申请接口时，如果协议类型为11、12、14、15时不可为空
	BankType string
	//卡号 可空(本次交易中,从客户的哪张卡上扣钱) 协议类型：11：借记卡扣款 12：信用卡扣款 13：余额扣款 14：	余额+借记卡扣款15： 余额+信用卡扣款协议代扣申请接口时，如果协议类型为11、12、14、15时不可为空
	BankCardNo string
	//扩展字段 可空 当商户为二级商户是此字段必填  secMerchantId^二级商户ID|mccCode^mcc码
	ExtendParams string
	//商户银行账户 可空 商户用于收款的银行账户 如果资金到账T+0模式时必填。
	MerchantBankAccountNo string
	//商户手续费 本次交易需要扣除的手续费。单位:分,不支持小数点
	RateAmount string
	//客户信用卡有效期 可空 如果协议类型为12、15时不可为空
	CustCardValidDate string
	//客户信用卡的cvv2 可空 如果协议类型为12、15时不可为空
	CustCardCvv2 string
	//商户后台通知URL 当交易完成后，快付通会URL异步通知商户(只有非终态才有回调) 回调请求参数：orderNo=9897013867 amount=100 merchantId=2018041100 status=1 failureDetails=成功 errorCode=xxxx
	NotifyUrl string
	//城市编码 二级商户账户编号为空时，根据城市代码选择匹配的二级商户账户编号
	CityCode string
	//公网IP地址 小额业务必传
	SourceIP string
	//设备标识 移动终端唯一标识
	DeviceID string
}

/*
*3.8快捷协议代扣接口返回对象
 */
type Return_38 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.8快捷协议代扣接口异步回传对象
 */
type Return_38_A struct {
	//语言
	Language string `json:"language"`
	//调用端IP
	CallerIp string `json:"callerIp"`
	//参数签名算法
	SignatureAlgorithm string `json:"signatureAlgorithm"`
	//签名值
	SignatureInfo string `json:"signatureInfo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//交易金额
	Amount string `json:"amount"`
	//商户编号
	MerchantId string `json:"merchantId"`
}

/*
*3.9快捷协议代扣协议查询接口
 */
type Model_39 struct {
	//订单编号
	OrderNo string
	//协议号
	TreatyNo string
}

/*
*3.9快捷协议代扣协议查询接口返回对象
 */
type Return_39 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//协议号
	TreatyNo string `json:"treatyNo"`
	//协议类型
	TreatyType string `json:"treatyType"`
	//交易状态
	Status string `json:"status"`
	//协议生效日期
	StartDate string `json:"startDate"`
	//协议失效日期
	EndDate string `json:"endDate"`
	//协议相关说明
	Note string `json:"note"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.10快捷协议代扣协议解除接口
 */
type Model_310 struct {
	//订单编号
	OrderNo string
	//协议号
	TreatyNo string
}

/*
*3.10快捷协议代扣协议解除接口返回对象
 */
type Return_310 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//协议号
	TreatyNo string `json:"treatyNo"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.11用户资金查询接口
 */
type Model_311 struct {
	//请求编号
	ReqNo string
	//身份证号码
	CustID string
	//页码
	PageNum string
}

/*
*3.11用户资金查询接口返回对象
 */
type Return_311 struct {
	//请求编号
	ReqNo string `json:"seqNo"`
	//交易状态 1:成功 2:失败
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//列表详情
	Details []Return_311_Details `json:"details"`
	//总共页数
	TotalPage string `json:"totalPage"`
	//当前页码
	PageNum string `json:"pageNum"`
	//每页记录数
	PageSize string `json:"pageSize"`
	//总记录数
	TotalCount string `json:"totalCount"`
}

/*
*3.11用户资金查询接口返回详细对象
 */
type Return_311_Details struct {
	BalanceAmount string `json:"balanceAmount"`
	BalanceDate   string `json:"balanceDate"`
	CustID        string `json:"custID"`
	MerchantId    string `json:"merchantId"`
}

/*
*3.12交易查询接口接口
 */
type Model_312 struct {
	//开始日期
	StartDate string
	//结束日期
	EndDate string
	//商户订单号 可空
	OrderNo string
	//交易类型 可空
	TradeType string
	//交易状态 可空
	Status string
}

/*
*3.12交易查询接口返回对象
 */
type Return_312 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//交易状态 1:成功 2:失败
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//列表详情
	Details interface{} `json:"details"`
}

/*
*3.13账户余额查询接口
 */
type Model_313 struct {
}

/*
*3.13账户余额查询接口返回对象
 */
type Return_313 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//账户余额
	AvailableBalance string `json:"availableBalance"`
}

/*
*3.14交易类对账文件获取接口
 */
type Model_314 struct {
	CheckDate string
}

/*
*3.14交易类对账文件获取接口返回对象
 */
type Return_314 struct {
	//总笔数
	TotalCount string `json:"totalCount"`
	//总金额
	TotalAmount string `json:"totalAmount"`
	//交易状态
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//列表详情
	Details interface{} `json:"details"`
}

/*
*3.15银行卡三要素验证接口
 */
type Model_315 struct {
	//订单编号 系统自动生成
	//OrderNo string
	//客户银行卡行别
	CustBankNo string
	//客户姓名
	CustName string
	//客户银行账户号
	CustBankAccountNo string
	//客户账户借记贷记类型
	CustAccountCreditOrDebit string
	//客户证件类型
	CustCertificationType string
	//客户证件号码
	CustID string
}

/*
*3.15银行卡三要素验证接口返回对象
 */
type Return_315 struct {
	//请求编号
	ReqNo string `json:"reqNo"`
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*3.16验证类对账文件获取接口
 */
type Model_316 struct {
	//对账日期
	CheckDate string
}

/*
*3.16验证类对账文件获取接口返回对象
 */
type Return_316 struct {
	//总笔数
	TotalCount string `json:"totalCount"`
	//总金额
	TotalAmount string `json:"totalAmount"`
	//交易状态
	Status string `json:"status"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
	//列表详情
	Details interface{} `json:"details"`
}

/*
*3.17验证类对账文件获取接口返回对象
 */
type Return_317 struct {
	//订单编号
	OrderNo string `json:"orderNo"`
	//交易状态
	Status string `json:"status"`
	//银行返回时间
	BankReturnTime string `json:"bankReturnTime"`
	//失败详情
	FailureDetails string `json:"failureDetails"`
	//错误码
	ErrorCode string `json:"errorCode"`
}

/*
*结算3.22查询账户余额接口返回对象
 */
type Return_322 struct {
	//账户余额
	AvailableBalance int `json:"availableBalance"`
}
type KFT_BindCardConfirmInput struct {
	HolderName string
	//默认""
	BankCardNo  string
	OrderNumber string
	//短信流水
	SmsSeq string
	//默认""
	AuthCode string
	//默认"2"
	BankCardType string
	//默认"0321"
	CustCardValidDate string
	//默认"632"
	CustCardCvv2 string
	UserId       string
}

type KFT_BindCardInput struct {
	//默认"秦鹏超"
	HolderName string
	//默认"3011000"
	BankType string
	//默认"2"
	BankCardType string
	//默认"0321"
	CustCardValidDate string
	//默认"632"
	CustCardCvv2 string
	//默认"6222520626626549"
	BankCardNo string
	//默认"15890102164"
	MobileNo string
	//默认"410521198804150138"
	CertificateNo string
	UserId        string
}

type KFT_DeletaPayInput struct {
	R1_merchantNo string
	OrderNumber   string
	TreatyNo      string
	Amount        string
	RateAmount    string
	HolderName    string
	//默认""
	BankCardNo string
	//默认""
	BankType string
	//默认""
	CustCardValidDate string
	//默认""
	CustCardCvv2 string
	//默认0
	CardType int
}

type KFT_QuickPayQueryInput struct {
	OrderNumber string
	StartDate   string
	EndDate     string
}

type KFT_ThreeVerification struct {
	CustBankNo               string
	CustName                 string
	CustBankAccountNo        string
	CustAccountCreditOrDebit string
	CustID                   string
}

/*
结果信息
*/
type ResultMsg struct {
	//信息状态
	Success bool
	//信息内容
	Info string
}

/*----------------------------绑卡相关信息实体开始------------------------------------------*/
/*
银行卡基础信息
*/
type CreditBindInput struct {
	//用户编号
	UserId string
	//银行卡
	CardNo string
	//预留手机号
	Phone string
	//有效期
	ExpireTime string
	//安全码
	Cvv2 string
	//类型
	BankCardType string
}

/*
银行卡扩展信息
*/
type CreditBindInputChannel struct {
	CreditBindInput
	//银行卡
	BankId string
	//通道
	DefaultId string
}

/*----------------------------绑卡相关信息实体结束------------------------------------------*/
