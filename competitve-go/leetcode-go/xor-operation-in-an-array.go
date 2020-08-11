func xorOperation(n int, start int) int {
    temp := start
    for i:=1; i<n; i++ {
        temp = temp ^ (start+2*i)
    }
    return temp
}