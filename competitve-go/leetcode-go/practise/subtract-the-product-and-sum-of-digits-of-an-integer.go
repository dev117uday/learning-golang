func subtractProductAndSum(n int) int {
    result := []int{1,0}
    for n>0 {
        temp := n%10
        result[0] = result[0]*temp
        result[1] = result[1]+temp
        n = n/10
    }
    return result[0]-result[1]
}