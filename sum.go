package pairsum

func DoesItSum(a []int, sum int) bool {
	left := 0
	right := len(a) - 1
	for left < right {
		s := a[left] + a[right]
		if s == sum {
			return true
		}
		if s < sum {
			left++
			continue
		}
		right--
	}
	return false
}
