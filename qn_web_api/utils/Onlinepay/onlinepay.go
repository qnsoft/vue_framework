package onlinepay

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"qnsoft/qn_web_api/utils/DateHelper"
	"qnsoft/qn_web_api/utils/ErrorHelper"
	"qnsoft/qn_web_api/utils/php2go"
)

/*
*各种资金包生成
 */
//func main() {
//随机红包生成设置
//	redPackage(10, 200)
//随机代付包生成设置
//	DFPackage(3, 10000)
//}
/*------------红包开始--------------*/
// 随机红包
// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}

	rand.Seed(time.Now().UnixNano())

	var min = 1
	max := remainMoney / remainCount * 2
	money := rand.Intn(max) + min
	return money
}

// 发红包
// count: 红包数量
// money: 红包金额（单位：分)
func redPackage(count, money int) {
	for i := 0; i < count; i++ {
		m := randomMoney(count-i, money)
		//	fmt.Printf("%d  ", m)
		money -= m
	}
}

/*------------红包结束--------------*/
/*------------代付包开始--------------*/
// 随机代付包
// remainCount: 剩余包数
// remainMoney: 剩余包金额（单位：分)
func DFrandomMoney(remainCount, remainMoney, min, ylj int) int {
	if remainCount == 1 {
		return remainMoney / remainCount
	}
	//var min = canshu
	max := remainMoney / remainCount
	money := max - rand.Intn(min)
	if money > ylj {
		money = ylj
	}
	return money
}

// 发代付包
// count: 包数量
// money: 包金额（单位：分)
//ylj: 预留金（单位：分)
func DFPackage(count, money, canshu, ylj int) []int {
	var rt []int
	for i := 0; i < count; i++ {
		m := DFrandomMoney(count-i, money, canshu, ylj)
		fmt.Printf("%d  ", m)
		rt = append(rt, m)
		money -= m
	}
	return rt
}

/*-----------代付包结束---------------*/
/*------------时间包开始--------------*/
// 时间分发包(日)
// count: 包数量
// satart:开始时间
// end: 结束时间
func PackageDay(count int, days string) ([]string, []int, int) {
	var _rt []int
	var _str []string
	str_days := php2go.Rtrim(days, "|")
	arry_adys := strings.Split(str_days, "|")
	for i := 0; i < len(arry_adys); i++ {
		_Time, err := date.ParseLocal(arry_adys[i])
		ErrorHelper.CheckErr(err)
		for j := 0; j < count; j++ {
			_rt = append(_rt, _Time.Day())
			//_rt = append(_rt, _Time.Day())
			_str = append(_str, arry_adys[i])
		}
	}
	return _str, _rt, len(arry_adys) * count
}

// 时间分发包(日)
// count: 包数量
// satart:开始时间
// end: 结束时间
func PackageHour(count, day_count, satart, end int) []int {
	var _rt []int
	if day_count == 1 {
		for i := 1; i <= count; i++ {
			num := php2go.Rand(satart, end)
			_rt = append(_rt, num)
		}
	} else if day_count == 2 {
		for i := 1; i <= count; i++ {
			num := 0
			for j := 1; j <= 2; j++ {
				if j == 1 {
					num = php2go.Rand(8, 11)
					_rt = append(_rt, num)
				} else if j == 2 {
					num = php2go.Rand(13, 19)
					_rt = append(_rt, num)
				}
			}
		}
	} else if day_count == 3 {
		for i := 1; i <= count; i++ {
			num := 0
			for j := 1; j <= 3; j++ {
				if j == 1 {
					num = php2go.Rand(8, 11)
					_rt = append(_rt, num)
				} else if j == 2 {
					num = php2go.Rand(12, 15)
					_rt = append(_rt, num)
				} else if j == 3 {
					num = php2go.Rand(16, 19)
					_rt = append(_rt, num)
				}
			}
		}
	} else {
		for i := 1; i <= count; i++ {
			num := 0
			for j := 1; j <= 3; j++ {
				if j == 1 {
					num = php2go.Rand(8, 11)
					_rt = append(_rt, num)
				} else if j == 2 {
					num = php2go.Rand(12, 15)
					_rt = append(_rt, num)
				} else if j == 3 {
					num = php2go.Rand(16, 19)
					_rt = append(_rt, num)
				}
			}
		}
	}
	return _rt
}

// 时间分发包
// count: 包数量
// satart:开始时间
// end: 结束时间
func PackageTime(count, satart, end int) []int {
	var _rt []int
	for i := 1; i <= count; i++ {
		now := 0
		if satart == end {
			now = satart
			_rt = append(_rt, now)
		} else {

			now = php2go.Rand(satart, end-2)
			if now+i < 58 {
				_rt = append(_rt, now+i)
			} else {
				_rt = append(_rt, i)
			}
		}

	}
	//_rt = Mp_Paixu(_rt)
	return _rt
}

/*
*冒泡排序
 */
func Mp_Paixu(_arry []int) []int {
	num := len(_arry)
	for i := num; i > 0; i-- {
		for j := i + 1; j < num; j++ {
			if _arry[i] < _arry[j] {
				tmp := _arry[i]
				_arry[i] = _arry[j]
				_arry[j] = tmp
			}
		}
	}
	return _arry
}

/*-----------代付包结束---------------*/
