func defangIPaddr(address string) string {
    var result string
    for i:=0; i<len(address); i++ {
        if string(address[i]) == "." {
            result = result + "[.]"
        } else {
            result = result + string(address[i])
        }
    }
    return result
}