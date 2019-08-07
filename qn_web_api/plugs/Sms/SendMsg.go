package Sms

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//短信内容实体
type Sms_Model struct {
	//用户名-APIID
	Account string
	//密码-APIKEY
	Password string
	//手机号
	Mobile string
	//内容
	Content string
}

//短信接口
type SendSms interface {
	SendMessage() Sms_Model
}

//互亿无线
type HyWx struct {
	//Content Sms_Model
}

//返回值信息实体
type Sms_rtjson_Model struct {
	//状态码
	Code int `json:"code"`
	//消息内容
	Msg string `json:"msg"`
	//消息id
	Smsid string `json:"smsid"`
}

func (h HyWx) SendMessage() Sms_Model {
	return Sms_Model{}
}

/*
//验证码发送
@param _s
*/
func SendMsg(_s SendSms, _model Sms_Model) string {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	v.Set("account", _model.Account)                                                                  //_model.Account)                                          //用户名是登录用户中心->验证码短信->产品总览->APIID
	v.Set("password", GetMd5String(_model.Account+_model.Password+_model.Mobile+_model.Content+_now)) //查看密码请登录用户中心->验证码短信->产品总览->APIKEY
	v.Set("mobile", _model.Mobile)                                                                    //_model.Mobile)                        //手机号 "136xxxxxxxx"
	v.Set("content", _model.Content)                                                                  // _model.Content)                      //内容 "您的验证码是：9552。请不要把验证码泄露给其他人。"
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构
	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return string(data)
}

//MD5加密
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//Unicode解密
func u2s(form string) (to string, err error) {
	bs, err := hex.DecodeString(strings.Replace(form, `\u`, ``, -1))
	if err != nil {
		return "", err
	}
	for i, bl, br, r := 0, len(bs), bytes.NewReader(bs), uint16(0); i < bl; i += 2 {
		binary.Read(br, binary.BigEndian, &r)
		to += string(r)
	}
	return to, err
}

//发送短信验证码
func SenMsg() {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := "用户名" //用户名是登录用户中心->验证码短信->产品总览->APIID
	_password := "密码" //查看密码请登录用户中心->验证码短信->产品总览->APIKEY
	_mobile := "136xxxxxxxx"
	_content := "您的验证码是：9552。请不要把验证码泄露给其他人。"
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}
