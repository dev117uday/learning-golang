func removeElement(nums []int, val int) int {
    answer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			copy(nums[i:], nums[i+1:])
			nums[len(nums)-1] = 0
			nums = nums[:len(nums)-1]
			answer++
			i--
		}
	}

	return nums
}
