package rand_to_rand

import (
	"math/rand"
	"time"
)

// F 此函数只能用，不能修改
// 等概率返回1~5之间的随机整数
func F() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(5) + 1
}

// A 等概率得到0和1两个数
func A() int {
	ans := 0
	for {
		ans = F()
		if ans != 3 {
			break
		}
	}
	if ans < 3 {
		return 0
	} else {
		return 1
	}
}

// B 等概率返回0~6之间的随机整数
func B() int {
	ans := 0
	for {
		ans = (A() << 2) + (A() << 1) + A()
		if ans != 7 {
			break
		}
	}
	return ans
}

// C 等概率返回1~7之间的随机整数
func C() int {
	return B() + 1
}
