func numIdenticalPairs(nums []int) int {
    goodPairs := make(map[int]int)
    result := 0

    for _, num := range nums {
        result += goodPairs[num]
        goodPairs[num]++
    }
    return result
}
