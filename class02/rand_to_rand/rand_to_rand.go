package rand_to_rand

import (
	"math"
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

// X 你只能知道，X会以固定概率返回0和1，但是X的具体实现是什么，你不知道。
func X() int {
	randomNum := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	if randomNum < 0.84 {
		return 0
	}
	return 1
}

// Y 等概率返回0和1
func Y() int {
	ans := X()
	for ans == X() {
	}
	return ans
}

// XToPower2
// 返回[0,1)的一个小数
// 任意的x，x属于[0,1)，[0,x)范围上的数出现概率由原来的x调整成x平方
func XToPower2() float64 {
	return math.Min(
		rand.New(rand.NewSource(time.Now().UnixNano())).Float64(),
		rand.New(rand.NewSource(time.Now().UnixNano())).Float64(),
	)
}

// XToPower3
// 返回[0,1)的一个小数
// 任意的x，x属于[0,1)，[0,x)范围上的数出现概率由原来的x调整成x立方
func XToPower3() float64 {
	return math.Min(
		rand.New(rand.NewSource(time.Now().UnixNano())).Float64(),
		math.Min(
			rand.New(rand.NewSource(time.Now().UnixNano())).Float64(),
			rand.New(rand.NewSource(time.Now().UnixNano())).Float64(),
		),
	)
}
