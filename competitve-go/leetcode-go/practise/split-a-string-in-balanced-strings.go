func balancedStringSplit(s string) int {
    count := 0
    lCount := 0
    rCount :=0
    for i:=0; i<len(s); i++ {
        if string(s[i]) == string("L") {
            lCount++
        } else {
            rCount++
        }
        if lCount == rCount {
            count++
            lCount = 0
            rCount = 0
        }
    }
    return count
}