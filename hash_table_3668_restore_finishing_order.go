func recoverOrder(order []int, friends []int) []int {
    frequency := make(map[int]int)

    for _, o := range order {
        frequency[o]++
    }

    for _, f := range friends {
        frequency[f]++
    }
    result := []int{}
    for _, o := range order {
        val, ok := frequency[o]
        if ok && val > 1 {
            result = append(result, o)
        }
    }
    return result
}