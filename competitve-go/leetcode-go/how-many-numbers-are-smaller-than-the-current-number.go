func smallerNumbersThanCurrent(nums []int) []int {
    var result []int
    for i:=0; i<len(nums); i++ {
        count := 0
        for j:=0; j<len(nums); j++ {
            if nums[j] < nums[i] && j!=i {
                count++
            }
        }
        result = append(result,count)
    }
    return result
}