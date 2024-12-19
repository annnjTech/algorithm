package bit_calculator

import "math"

/**
都是整数操作
*/

// Add 计算两个整数的相加
func Add(a, b int64) int64 {
	sum := a
	for b != 0 { // 直到b这个进位信息为00000000结束循环，跳出循环时候，a和sum一致
		// a=46 b=20 a+b=66
		/*
			异或=不进位相加
			    a = 00101110
			    b = 00010100
			a ^ b = 00111010
		*/
		sum = a ^ b
		/*
			           a = 00101110
				       b = 00010100
			       a & b = 00000100
			(a & b) << 1 = 00001000 //这里的意思是，a和b与完的结果向左移动一位，结果就是a加b的进位信息
		*/
		b = (a & b) << 1 // 将进位信息赋值给b
		a = sum          // 此时的无进位相加的结果，赋值给a，继续下一次循环
	}
	return sum
}

// NegNum 计算一个整数的相反数，即取反加1
func NegNum(n int64) int64 {
	return Add(^n, 1)
}

// Minus 计算两个整数的相减
func Minus(a, b int64) int64 {
	return Add(a, NegNum(b)) // a-b相当于a+(-b)
}

// Multi 计算两个整数的相乘
func Multi(a, b int64) int64 {
	var res int64
	for b != 0 {
		if (b & 1) != 0 {
			res = Add(res, a)
		}
		a <<= 1
		b = int64(uint64(b) >> 1)
	}
	/*
		初始：
		a = 00000110 = 6
		b = 00000101 = 5
		第一次：
		b & 1 = 00000101 & 00000001 != 0，走res加的逻辑
		res = 00000000 + 00000110 = 00000110 = 6
		a = 00001100
		b = 00000010
		第二次：
		b & 1 = 00000010 & 00000001 = 0，不走res加的逻辑
		a = 00011000
		b = 00000001
		第三次：
		b & 1 = 00000001 & 00000001 = 1，走res加的逻辑
		res = 00000110 + 00011000 = 6 + 24 = 30
		a = 00110000
		b = 00000000
		第四次：
		b != 0 为false，结束循环
	*/
	return res
}

// IsNeg 判断一个整数是否为负数
func IsNeg(n int64) bool {
	return n < 0
}

func div(a, b int64) int64 {
	var res int64
	x, y := a, b
	// 需要将a，b转化为正数
	if IsNeg(a) {
		x = NegNum(a)
	}
	if IsNeg(b) {
		y = NegNum(b)
	}

	/*
			x: 00110000 = 48
			y: 00000101 = 5，x/y的结果应该是9
			循环：
			x 右移 30位 = 0.....00000000，>= y的00000101吗？不大于等于，继续下一次循环
			x 右移 29位 = 0.....00000000，>= y的00000101吗？不大于等于，继续下一次循环
			x 右移 28位 = 0.....00000000，>= y的00000101吗？不大于等于，继续下一次循环
			...
			x 右移 4位 = 0.....00000011，>= y的00000101吗？不大于等于，继续下一次循环
			x 右移 3位 = 0.....00000110，>= y的00000101吗？大于等于，进入if，res |= 1左移3位，res = 0...00000000 | 0...00001000 = 0...00001000
			   x = 00110000
			y<<3 = 00101000 ,相减后，x = 00001000
			x 右移 2位 = 0.....00000010，>= y的00000101吗？不大于等于，继续下一次循环
			x 右移 1位 = 0.....00000100，>= y的00000101吗？不大于等于，继续下一次循环
			x 右移 0位 = 0.....00001000，>= y的00000101吗？大于等于，进入if，res |= 1左移0位，res = 0...00001000 | 0...00000001 = 0...00001001
			i = -1，结束循环
		    此时的结果res = 00001001 = 9，即x/y的结果
	*/
	var i int64
	for i = 30; i >= 0; i = Minus(i, 1) {
		if (x >> i) >= y {
			res |= 1 << i
			x = Minus(x, y<<i)
		}
	}
	if (!IsNeg(a) && IsNeg(b)) || (IsNeg(a) && !IsNeg(b)) {
		res = NegNum(res)
	}
	return res
}

func Divide(a, b int64) int64 { //a是除数，b是被除数
	if a == math.MinInt64 && b == math.MinInt64 {
		return 1
	} else if b == math.MinInt64 {
		return 0
	} else if a == math.MinInt64 {
		if b == NegNum(1) {
			return math.MaxInt64
		} else {
			var c int64
			c = div(Add(a, 1), b)
			return Add(c, div(Minus(a, Multi(c, b)), b))
		}
	} else {
		return div(a, b)
	}
}
