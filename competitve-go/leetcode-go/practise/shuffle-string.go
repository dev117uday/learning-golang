func restoreString(s string, indices []int) string {
    var result string = ""
    for i:=0; i<len(indices); i++ {
        for j:=0; j<len(indices); j++ {
            if i == indices[j]{
                result = result + string(s[j])
                temp := indices
                var indices []int
                for x:=0; x<=j-1; x++ {
                    indices = append(indices,temp[x])
                } 
                for x:=j+1; x<len(temp); x++ {
                    indices = append(indices,temp[x])
                } 
                break
            }

        }
    }
    return result
}