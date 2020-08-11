func runningSum(nums []int) []int {
    var result []int
    sum := 0
    for _,value := range nums{
        sum = sum + value
        result = append(result,sum)
    }
    return result
}