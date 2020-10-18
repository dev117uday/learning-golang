func findMaxConsecutiveOnes(nums []int) int {
    counter := 0
    max := 0 
    for i:=0; i<len(nums); i++ {
        if nums[i] == 1 {
            counter++
        } else {
            if counter > max {
                max = counter
            }
            counter = 0
        }
    }
    if counter > max {
       return counter
    } else {
        return max
    }
}
