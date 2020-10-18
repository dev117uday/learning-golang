func findNumbers(nums []int) int {
 
    answer := 0
    for i:=0;i<len(nums);i++ {
        if numberOfDigits(nums[i])%2 == 0 {
            answer++
        }
    }
    return answer
}

func numberOfDigits (num int) int {
    count := 0
    for num > 0 {
        num = num/10
        count++
    }
    return count
}
