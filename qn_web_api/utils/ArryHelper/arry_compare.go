package arry

/*
获取int数组中最大的值和下标
*/
func Max_IntArry_One(_arry []int) (int, int) {
	//获取一个数组里最大值，并且拿到下标

	//声明一个数组5个元素
	// var arr [5]int = [...]int {6, 45, 63, 16 ,86}
	//假设第一个元素是最大值，下标为0
	maxVal := _arry[0]
	maxIndex := 0

	for i := 1; i < len(_arry); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < _arry[i] {
			maxVal = _arry[i]
			maxIndex = i
		}
	}
	return maxVal, maxIndex
}

// /**
// 根据key排序
// */
// func sortMap2(mp map[string]string) map[string]string {
// 	var newMp = make([]string, 0)
// 	for k, _ := range mp {
// 		newMp = append(newMp, k)
// 	}
// 	sort.Strings(newMp)
// 	for _, v := range newMp {
// 		fmt.Println("根据key排序后的新集合》》   key:", v, "    value:", mp[v])
// 	}
// 	return newMp
// }

// /**
// 根据value排序
// */
// func sortMap(mp map[string]string) map[string]string {
// 	var newMp = make([]int, 0)
// 	var newMpKey = make([]string, 0)
// 	for oldk, v := range mp {
// 		newMp = append(newMp, v)
// 		newMpKey = append(newMpKey, oldk)
// 	}
// 	sort.Ints(newMp)
// 	for k, v := range newMp {
// 		fmt.Println("根据value排序后的新集合》》  key:", newMpKey[k], "    value:", v)
// 	}
// 	return newMp
// }
