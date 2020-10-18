func kidsWithCandies(candies []int, extraCandies int) []bool {
    var result []bool
    biggest := candies[0]
    for _,value := range candies {
        if value > biggest {
            biggest = value
        }
    }
    
    for _,value := range candies {
        if extraCandies + value >= biggest {
            result = append(result,true)
        } else {
            result = append(result,false)
        }
    }
    return result
}