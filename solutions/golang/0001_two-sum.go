func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		her := target - nums[i]
		if herIndex, ok := m[her]; ok {
			return []int{herIndex, i}
		}
		m[nums[i]] = i
	}
	return nil
}
