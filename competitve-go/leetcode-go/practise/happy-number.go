func isHappy(n int) bool {
	var index int = 0
	set := make(map[int]bool)
	set[n] = true
	for {
		var temp int = 0
		for n != 0 {
			index = n % 10
			temp = temp + index*index
			n = n / 10
		}

		if temp == 1 {
			return true
		}

		_, answer := set[temp]
		if answer == true {
			return false
		} else {
			set[temp] = true
			n = temp
		}
	}
}