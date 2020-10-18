func shuffle(nums []int, n int) []int {
    var result []int
    length := len(nums)
    for i:=0; i<length/2; i++ {
        result = append(result,nums[i])
        result = append(result,nums[length/2+i])
    }
    return result
}