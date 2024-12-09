package factorial

/**
 * 求1!+2!+3!+...+n!
 */

// F1 分别算出1! 2! 3! ... n! 然后相加
func F1(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += factorial(i)
	}
	return
}

func factorial(n int) (sum int) {
	sum = 1
	for i := 1; i <= n; i++ {
		sum *= i
	}
	return
}

// F2 利用累积和公式求和
func F2(n int) (sum int) {
	current := 1
	for i := 1; i <= n; i++ {
		current *= i
		sum += current
	}
	return
}
