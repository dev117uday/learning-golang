// wrong method
func merge(nums1 []int, m int, nums2 []int, n int) {

	if len(nums2) == 0 {

	} else if m == 0 {
		copy(nums1[0:], nums2[:])
	} else {
		copy(nums1[m:], nums2[0:])

		swapped := true
		for swapped {
			swapped = false
			for i := 1; i < len(nums1); i++ {
				if nums1[i-1] > nums1[i] {
					nums1[i], nums1[i-1] = nums1[i-1], nums1[i]
					swapped = true
				}
			}
		}

	}

}

