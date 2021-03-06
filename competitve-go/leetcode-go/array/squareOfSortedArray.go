func sortedSquares(A []int) []int {
    var answer []int
    for _,val := range A {
        answer = append(answer, val*val)
    }
     return bubblesort(answer)
}

func bubblesort(items []int) []int {
    var (
        n = len(items)
        sorted = false
    )
    for !sorted {
        swapped := false
        for i := 0; i < n-1; i++ {
            if items[i] > items[i+1] {
                items[i+1], items[i] = items[i], items[i+1]
                swapped = true
            }
        }
        if !swapped {
            sorted = true
        }
        n = n - 1
    }
    return items
}
