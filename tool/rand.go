package tool

import (
	"math/rand"
	"time"
)

// RandNumStr 生成指定位数的随机数，不能太大
func RandNumStr(length int) string {
	if length <= 0 || length > 10 {
		length = 10
	}

	strNums := "1234567890"
	var num string
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		num = num + string(strNums[rand.Intn(len(strNums))])
	}
	return num
}

func RandStr(length int) string {
	if length <= 0 || length > 20 {
		length = 20
	}

	str := "qwertyuiopasdfghjklzxcvbnm"
	var s string
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		s = s + string(str[rand.Intn(len(str))])
	}
	return s
}
