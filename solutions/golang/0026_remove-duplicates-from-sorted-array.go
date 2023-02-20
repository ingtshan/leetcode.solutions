func removeDuplicates(nums []int) int {
	cp, mp := 0, 1
	for ; mp < len(nums); mp++ {
		if nums[cp] != nums[mp] {
			cp++
			nums[cp] = nums[mp]
		}
	}
	return cp + 1
}
