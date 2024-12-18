package comp

func TestBsExists(sortedArr []int, num int) bool {
	for _, curr := range sortedArr {
		if curr == num {
			return true
		}
	}
	return false
}
