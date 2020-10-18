func removeDuplicates(nums []int) int {
    for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			copy(nums[i:], nums[i+1:])
			nums = nums[:len(nums)-1]
			i--
		}
	}
    return len(nums)
}
