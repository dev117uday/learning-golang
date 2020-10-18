func createTargetArray(nums []int, index []int) []int {
    if len(index) == 1 {
        return nums
    }
    
    var result []int
    for i:=0; i<len(nums); i++ {
        if index[i] >= i {
            result = append(result,nums[i])
        } else {
            
            var temp []int
            for j:=0; j<index[i]; j++ {
                temp = append(temp,result[j])
            }
            temp = append(temp,nums[i])
            for j:=index[i]; j<len(result); j++ {
                temp = append(temp,result[j])
            }
            result = result[:0]
            for j:=0; j<len(temp); j++ {
                result = append(result,temp[j])
            }
             
        }
    }
    return result
}