package max_min

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	mx := nums[0]
	for _, x := range nums {
		if x > mx {
			mx = x
		}
	}
	return mx
}
func min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	mi := nums[0]
	for _, x := range nums {
		if x < mi {
			mi = x
		}
	}
	return mi
}
