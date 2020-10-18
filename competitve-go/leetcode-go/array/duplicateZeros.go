func duplicateZeros(arr []int)  {
	var lenArr = len(arr) - 1

	for i := 0; i < lenArr; i++ {
		if arr[i] == 0 {
			copy(arr[i+1:], arr[i:])
			arr[i] = 0
			i++
		}

	}
}
