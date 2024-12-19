package main

import (
	"algorithm/class05/bit_calculator"
	"fmt"
	"math"
)

func main() {
	fmt.Println(bit_calculator.Add(46, 20))
	fmt.Println(bit_calculator.NegNum(-46))
	fmt.Println(bit_calculator.Minus(2, 1))
	fmt.Println(bit_calculator.Multi(8, 3))
	fmt.Println(bit_calculator.Divide(math.MinInt64, math.MinInt64))
	fmt.Println(bit_calculator.Divide(1, math.MinInt64))
	fmt.Println(bit_calculator.Divide(math.MinInt64, -1))
	fmt.Println(math.MaxInt64)
	fmt.Println(bit_calculator.Divide(math.MinInt64, 100))
	fmt.Println(bit_calculator.Divide(10, 3))
}
