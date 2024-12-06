package main

import (
	"algorithm/class02/rand_to_rand"
	"fmt"
)

func main() {
	//fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Float32())

	fmt.Println("测试开始")
	testTimes := 100000000
	//count := 0
	counts := make(map[int]int)
	for i := 0; i < testTimes; i++ {
		ans := rand_to_rand.F()
		counts[ans]++
	}
	for i := 0; i <= len(counts); i++ {
		fmt.Printf("%d这个数，出现了%d次\n", i, counts[i])
	}
	fmt.Println("===========")
}
