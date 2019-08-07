package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"os/exec"
	"runtime/debug"
	"strconv"
	"time"
	"qnsoft/qn_web_api/controllers/OnlinePay/Kuaifutong"
	"qnsoft/qn_web_api/models/shop"
	date "qnsoft/qn_web_api/utils/DateHelper"
	"qnsoft/qn_web_api/utils/DbHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/StringHelper"
	"qnsoft/qn_web_api/utils/WebHelper"
	"qnsoft/qn_web_api/utils/php2go"

	"github.com/astaxie/beego"
)

var mailTpl *template.Template

func init() {
	mailTpl, _ = template.New("mail_tpl").Parse(`
	你好 {{.username}}，<br/>

<p>以下是任务执行结果：</p>

<p>
任务ID：{{.task_id}}<br/>
任务名称：{{.task_name}}<br/>       
执行时间：{{.start_time}}<br />
执行耗时：{{.process_time}}秒<br />
执行状态：{{.status}}
</p>
<p>-------------以下是任务执行输出-------------</p>
<p>{{.output}}</p>
<p>
--------------------------------------------<br />
本邮件由系统自动发出，请勿回复<br />
如果要取消邮件通知，请登录到系统进行设置<br />
</p>
`)

}

type Job struct {
	id         int                                               // 任务ID
	logId      int64                                             // 日志记录ID
	name       string                                            // 任务名称
	task       *shop.UserCardJobOrder                            // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func NewJobFromTask(task *shop.UserCardJobOrder) (*Job, error) {
	if task.Id < 1 {
		return nil, fmt.Errorf("ToJob: 缺少id")
	}
	job := NewCommandJob(task.Id, task.TaskName, "task.Command")
	job.task = task
	//	job.Concurrent = task.Concurrent == 1
	return job, nil
}

func NewCommandJob(id int, name string, command string) *Job {
	job := &Job{
		id:   id,
		name: name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		cmd := exec.Command("/bin/bash", "-c", command)
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func (j *Job) Status() int {
	return j.status
}

func (j *Job) GetName() string {
	return j.name
}

func (j *Job) GetId() int {
	return j.id
}

func (j *Job) GetLogId() int64 {
	return j.logId
}

func (j *Job) Run() {
	if !j.Concurrent && j.status > 0 {
		beego.Warn(fmt.Sprintf("任务[%d]上一次执行尚未结束，本次被忽略。", j.id))
		return
	}

	defer func() {
		if err := recover(); err != nil {
			beego.Error(err, "\n", string(debug.Stack()))
		}
	}()

	if workPool != nil {
		workPool <- true
		defer func() {
			<-workPool
		}()
	}
	//开始执行计划任务 //根据任务id,从数据库表中取出要执行的任务
	beego.Debug(fmt.Sprintf("开始执行信用卡代还: %d", strconv.Itoa(j.id)+"["+j.name+"]"))

	j.status++
	defer func() {
		j.status--
	}()
	ZN_DF(j.id)
	fmt.Println("信用卡代还执行完毕!")
}

/*
*智能代付
 */
func ZN_DF(_id int) {
	//根据订单id获取当前要执行的订单
	_job_order_model := &shop.UserCardJobOrder{Id: _id}
	results, err := DbHelper.MySqlDb().Get(_job_order_model)
	ErrorHelper.CheckErr(err)
	if results {
		_SourceIP := Kuaifutong.Server_IP //公网IP地址 可空 小于等于1000该项必传
		_merchantNo := Kuaifutong.R1_merchantNo
		//_Amount_XXX := _job_order_model.Amount
		//ErrorHelper.CheckErr(err)
		//代扣协议id
		_Treatyid := ""
		/*小额临时注释 if _job_order_model.Amount <= 100000 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { */ //计算快付通小额费率
		if _job_order_model.Amount <= 10 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 {
			_Treatyid = Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall
			_SourceIP = Kuaifutong.Server_IP
			_merchantNo = Kuaifutong.R3_merchantNo
		} else { //如果没有小额签约先走大额通道
			_Treatyid = Get_Bank_Model(_job_order_model.Bankcardno).Treatyid
			_bank_code := Get_Bank_Model(_job_order_model.Bankcardno).BankCode //获取取银行行别代码
			fellv, err := strconv.ParseFloat(Get_Bank_T0_FL(_bank_code), 64)
			ErrorHelper.CheckErr(err)
			_job_order_model.Rateamounta = int(math.Floor(float64(_job_order_model.Amount)*fellv + 0.5))
		}
		_CityCode := ""                                 //城市编码 可空
		_DeviceID := ""                                 //设备标识 可空
		_NotifyUrl := Kuaifutong.QuickPay_Pay_notifyUrl //商户后台通知URL
		//开始执行代扣代还订单
		_model_38 := Kuaifutong.Model_38{
			OrderNo:   "DK" + date.FormatDate(time.Now(), "yyMMddHHmmss") + strconv.Itoa(php2go.Rand(10, 99)), //订单编号 自动生成
			TreatyNo:  _Treatyid,
			TradeTime: date.FormatDate(time.Now(), "yyyyMMddHHmmss"), //生成当前时间
			Amount:    strconv.Itoa(_job_order_model.Amount),         //代扣金额
			//CustAccountId: _CustAccountId,                              //账户ID 协议类型 11：借记卡扣款 12：信用卡扣款 13：余额扣款 余额+借记卡扣款15： 余额+信用卡扣款协议代扣申请接口时，如果协议类型为13、14、15时不可为空
			HolderName: Get_Bank_Model(_job_order_model.Bankcardno).Cardholder, //持卡人真实姓名
			BankType:   Get_Bank_Model(_job_order_model.Bankcardno).BankCode,   //银行行别
			BankCardNo: _job_order_model.Bankcardno,                            //银行卡号 卡1出金卡
			//	ExtendParams:          _ExtendParams,                    //扩展字段 当商户为二级商户是此字段必填
			MerchantBankAccountNo: _job_order_model.Merchantbankaccountno,                 //商户银行账户 卡2入金卡 从传参来 商户用于收款的银行账户 资金到账T+0模式时必填。
			RateAmount:            strconv.Itoa(_job_order_model.Rateamounta),             //商户手续费
			CustCardValidDate:     Get_Bank_Model(_job_order_model.Bankcardno).Expiretime, //客户信用卡有效期
			CustCardCvv2:          Get_Bank_Model(_job_order_model.Bankcardno).Cvv2,       //客户信用卡的cvv2
			NotifyUrl:             _NotifyUrl,                                             //商户后台通知URL 写死的常量直接调取
			CityCode:              _CityCode,                                              //城市编码 可空
			SourceIP:              _SourceIP,                                              //公网IP地址 可空 小于等于1000该项必传
			DeviceID:              _DeviceID,                                              //设备标识 可空
		}
		_kuaifutong := Kuaifutong.KuaiPayHelper{}
		_rerurn_38 := _kuaifutong.Gbp_same_id_credit_card_treaty_collect(_model_38, Get_Bank_Model(_job_order_model.Bankcardno).Treatytype, _merchantNo)
		_str_rt_38 := fmt.Sprintf("%+v", _rerurn_38)
		ErrorHelper.LogInfo("执行了吗？" + _str_rt_38)
		if _rerurn_38.Status != "" {
			//开始将代扣信息写入到信用卡订单表[lkt_user_card_order]
			_update_model_order := shop.UserCardJobOrder{
				UserId:   _job_order_model.UserId, //用户id
				Ordernoa: _rerurn_38.OrderNo,      //订单编号
				//Amount:                strconv.FormatFloat(float64(_Amount*100), 'f', 0, 64),     //订单金额
				//Rateamounta:           strconv.FormatFloat(float64(_RateAmount*100), 'f', 0, 64), //代扣手续费
				//Bankcardno:            _BankCardNo,                                               //出金卡号
				//Merchantbankaccountno: _MerchantBankAccountNo, //入金卡号
				ReturnA: _str_rt_38, //代扣返回信息
				AddTime: time.Now(), //订单创建时间
			}
			_count, err := DbHelper.MySqlDb().Id(_job_order_model.Id).Update(&_update_model_order)
			ErrorHelper.CheckErr(err)
			if _count > 0 {
				//开始利用3.11接口查询余额
				_model_311 := Kuaifutong.Model_311{
					CustID:  Get_Bank_Model(_job_order_model.Bankcardno).IdCard,
					PageNum: "1",
				}
				//调用快付通对象
				_rerurn_11 := Kuaifutong.Return_311{}
				//利用3.11查询余额
				//_kuaifutong := Kuaifutong.KuaiPayHelper{}
				/*小额临时注释 if _job_order_model.Amount <= 100000 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { //使用小额通道查询 */
				if _job_order_model.Amount <= 10 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { //使用小额通道查询
					_rerurn_11 = _kuaifutong.Gbp_same_id_credit_card_not_pay_balance_small(_model_311)
					//	_job_order_model.Rateamountb = _job_order_model.Rateamountb
					_job_order_model.Rateamountb = int(math.Floor(float64(_job_order_model.Amount-_job_order_model.Rateamounta)*0.0025 + 0.5))
					ErrorHelper.LogInfo("智能代还走小额出金后查询")
				} else {
					_rerurn_11 = _kuaifutong.Gbp_same_id_credit_card_not_pay_balance(_model_311)
					bank_code := Get_Bank_Model(_job_order_model.Bankcardno).BankCode //获取取银行行别代码
					fellv, err := strconv.ParseFloat(Get_ZFB_T0_FL(bank_code), 64)
					ErrorHelper.CheckErr(err)
					_job_order_model.Rateamountb = int(math.Floor(float64(_job_order_model.Amount-_job_order_model.Rateamounta)*fellv + 0.5))
					ErrorHelper.LogInfo("智能代还走大额出金后查询")
				}
				if len(_rerurn_11.Details) > 0 {
					//臻方便代付费率
					//_feilvZ_t0, err := strconv.ParseFloat(Get_ZFB_T0_FL(Get_Bank_Model(_job_order_model.Bankcardno).BankCode), 64) //出金卡费率
					//ErrorHelper.CheckErr(err)
					//开始处理未签约小额的订单
					/*
						if len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { //计算快付通小额费率
							_job_order_model.Rateamountb = _job_order_model.Rateamountb
							ErrorHelper.LogInfo("智能代还走小额出金后查询")
						} else { //如果没有签约先走大额通道 走固定费率
							bank_code := Get_Bank_Model(_job_order_model.Bankcardno).BankCode //获取取银行行别代码
							fellv, err := strconv.ParseFloat(Get_ZFB_T0_FL(bank_code), 64)
							ErrorHelper.CheckErr(err)
							_job_order_model.Rateamountb = int(math.Floor(float64(_job_order_model.Amount-_job_order_model.Rateamounta)*fellv + 0.5))
							ErrorHelper.LogInfo("智能代还走大额出金后查询")
						}
					*/
					//小额订单处理结束
					str_order := "DH" + date.FormatDate(time.Now(), "yyyyMMddHHmmss") + StringHelper.GetRandomNum(6)
					_model_35 := Kuaifutong.Model_35{
						OrderNo:   str_order, //订单编号 用于标识商户发起的一笔交易,在批量交易中,此编号可写在批量请求文件中,用于标识批量请求中的每一笔交易
						TradeName: "臻方便商城智能", //交易名称 由商户填写,简要概括此次交易的内容.用于在查询交易记录时,提醒用户此次交易具体做了什么事情
						//MerchantBankAccountNo:    "",                                            //商户银行账号 可空 商户用于付款的银行账户，资金到账T+0模式时必填。
						//MerchantBindPhoneNo:      "",                                            //商户开户时绑定的手机号（可空）对于有些银行账户被扣款时，需要提供此绑定手机号才能进行交易；此手机号和短信通知客户的手机号可以一致，也可以不一致
						Amount:                   strconv.Itoa(_job_order_model.Amount - _job_order_model.Rateamounta), //交易金额 此次交易的具体金额,单位:分,不支持小数点
						CustBankNo:               Get_Bank_Model(_job_order_model.Bankcardno).BankCode,                 //客户银行账户行别 客户银行账户所属的银行的编号,具体见第5.3.1章节
						CustBankAccountIssuerNo:  "",                                                                   //客户开户行网点号 可空 指支付系统里的行号，具体到某个支行（网点）号；
						CustBankAccountNo:        _job_order_model.Merchantbankaccountno,                               //客户银行账户号 本次交易中,往客户的哪张卡上付钱
						CustName:                 Get_Bank_Model(_job_order_model.Bankcardno).Cardholder,               //客户姓名 收钱的客户的真实姓名
						CustBankAcctType:         "",                                                                   //客户银行账户类型 可空 指客户的银行账户是个人账户还是企业账户
						CustAccountCreditOrDebit: "",                                                                   //客户账户借记贷记类型 可空 若是信用卡，则以下两个参数“信用卡有效期”和“信用卡cvv2”； 1借记 2贷记 4 未知
						CustCardValidDate:        "",                                                                   //客户信用卡有效期 可空 信用卡的正下方的四位数，前两位是月份，后两位是年份；
						CustCardCvv2:             "",                                                                   //客户信用卡的cvv2 可空 信用卡的背面的三位数
						CustID:                   Get_Bank_Model(_job_order_model.Bankcardno).IdCard,                   //客户证件号码
						CustPhone:                Get_Bank_Model(_job_order_model.Bankcardno).Mobile,                   //客户手机号 如果商户购买的产品中勾选了短信通知功能，则当商户填写了手机号时,快付通会在交易成功后通过短信通知客户
						Messages:                 "",                                                                   //发送客户短信内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
						CustEmail:                "",                                                                   //客户邮箱地址 可空 如果商户购买的产品中勾选了邮件通知功能，则当商户填写了邮箱地址时,快付通会在交易成功后通过邮件通知客户
						EmailMessages:            "",                                                                   //发送客户邮件内容 可空 如果填写了,则把此参数值的内容发送给客户；如果没填写，则按照快付通的默认模板发送给客户；
						Remark:                   "",                                                                   //备注 可空 商户可额外填写备注信息,此信息会传给银行,会在银行的账单信息中显示(具体如何显示取决于银行方,快付通不保证银行肯定能显示)
						CustProtocolNo:           Get_Bank_Model(_job_order_model.Bankcardno).Treatyid,                 //客户协议编号 可空 扣款人在快付通备案的协议号。
						ExtendParams:             "",                                                                   //扩展参数 可空 用于商户的特定业务信息传递，只有商户与快付通约定了传递此参数且约定了参数含义，此参数才有效。参数格式：参数名 1^参数值 1|参数名 2^参数值 2|……多条数据用“|”间隔注意: 不能包含特殊字符（如：#、%、&、+）、敏感词汇, 如果必须使用特殊字符,则需要自行做URL Encoding
						RateAmount:               strconv.Itoa(_job_order_model.Rateamountb),                           //商户手续费 可空 本次交易需要扣除的手续费。单位:分,不支持小数点 如不填，则手续费默认0元；
					}

					//_kuaifutong := Kuaifutong.KuaiPayHelper{}
					//开始利用35接口执行代付
					_return_35 := Kuaifutong.Return_35{}
					/*小额临时注释 if _job_order_model.Amount <= 100000 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { //处理小额代付 */
					if _job_order_model.Amount <= 10 && len(Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall) > 5 { //处理小额代付
						_model_35.CustProtocolNo = Get_Bank_Model(_job_order_model.Bankcardno).TreatyidSmall
						_return_35 = _kuaifutong.Gbp_same_id_credit_card_pay_small(_model_35)
						ErrorHelper.LogInfo("智能代还走小额入金")
					} else {
						_return_35 = _kuaifutong.Gbp_same_id_credit_card_pay(_model_35)
						ErrorHelper.LogInfo("智能代还走大额入金")
					}
					_str_rt_35 := fmt.Sprintf("%+v", _return_35)
					//_RateAmountAA, _ := strconv.Atoi(_RateAmountA)
					//开始更新数据库信用卡订单表[lkt_user_card_job_order]，补充代付部分订单信息
					_update_model_order := shop.UserCardJobOrder{
						UserId:   _job_order_model.UserId, //用户id
						Ordernob: _return_35.OrderNo,      //订单编号
						ReturnB:  _str_rt_35,              //代付返回信息
						IsFinish: 1,                       //订单完成状态
					}
					_count, err := DbHelper.MySqlDb().Id(_job_order_model.Id).Update(_update_model_order)
					ErrorHelper.CheckErr(err)
					//执行成功开始分润！
					FunRun(_job_order_model.UserId, strconv.Itoa(_job_order_model.ParentOrderId)+"|"+strconv.Itoa(_job_order_model.Id), strconv.Itoa(_job_order_model.Amount/100), _job_order_model.TaskName)
					fmt.Println(_count)
				}
			}

		}
	}
}

/*
*根据银行卡号获取获取绑卡信息
 */
func Get_Bank_Model(_Bank_No string) *shop.UserBankCard {
	_rt := &shop.UserBankCard{}
	_model_new := &shop.UserBankCard{BankCardNumber: _Bank_No}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		_rt = _model_new
	}
	return _rt
}

/*
*根据银行代码获取当前银行t+0费率
 */
func Get_Bank_T0_FL(_bank_code string) string {
	_rt := "0.0054"
	_model_new := &shop.ChannelBankKft{Bankcode: _bank_code}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		if _model_new.D0freerate != "" {
			_rt = _model_new.D0freerate
		}
	}
	return _rt
}

/*
*根据银行代码获取当前臻方便t+0费率
 */
func Get_ZFB_T0_FL(_bank_code string) string {
	_rt := "0.0010"
	_model_new := &shop.ChannelBankKft{Bankcode: _bank_code}
	has, err := DbHelper.MySqlDb().Get(_model_new)
	ErrorHelper.CheckErr(err)
	if has {
		if _model_new.D0myrate != "" {
			_rt = _model_new.D0myrate
		}
	}
	return _rt
}

/*
*分润处理
 */
func FunRun(user_id, order_no, amount, mark string) {
	type _rt struct {
		code int
		msg  string
	}
	_HeaderData := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	_BodyData := map[string]interface{}{"user_id": user_id, "order_no": order_no, "type": "2", "amount": amount, "mark": mark}
	_http_url := "https://shop.xhdncppf.com/index.php?module=app&action=index&store_id=8&app=calc_profit"
	_req := WebHelper.HttpPost(_http_url, _HeaderData, _BodyData)
	err := json.Unmarshal([]byte(_req), &_rt{})
	ErrorHelper.CheckErr(err)
}
