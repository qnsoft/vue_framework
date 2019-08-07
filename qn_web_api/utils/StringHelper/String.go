package StringHelper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//生成随机字符串(数字+大写小字母)
func GetRandomString(lens int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成随机字符串(数字)
func GetRandomNum(lens int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lens; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
*字符串左侧补中字符
@_old_str 原字符串
@_char要补的字符
@_len补后的长度
*/
func Str_Left(_old_str, _char string, _len int) string {
	_str := fmt.Sprintf("%"+_char+strconv.Itoa(_len)+"s", _old_str)
	return _str
}

/*
*获取中间字符串 GetBetweenStr("满500元减10元", "减", "元") 输出10
@_str 要处理的字符串
@start开始字符串
@end结束字符串
*/
func GetBetweenStr(_str, start, end string) string {
	n := strings.Index(_str, start)
	if n == -1 {
		n = 0
	} else {
		n = n + len(start) // 增加了else，不加的会把start带上
	}
	_str = string([]byte(_str)[n:])
	m := strings.Index(_str, end)
	if m == -1 {
		m = len(_str)
	}
	_str = string([]byte(_str)[:m])
	return _str
}
